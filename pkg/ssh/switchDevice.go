package ssh

import (
	"golang.org/x/crypto/ssh"
	"io"
	"sync"
	"time"
	"strings"
	"fmt"
)

// 这里的交换机设备代指 所有网络设备，例如：网关,AC,AP,交换机等

type SwitchClient struct {
	Client    *ssh.Client
	ClientConfig *ssh.ClientConfig
	Devices *SwitchDevices
	unusable bool
}

type SwitchDevices struct {
	Host     string // IP 地址
	Port     int	   // 端口 22
	UserName     string // 用户名
	Password string // 密码
	ProtocolType string //协议类型
}

func NewSwitchDevices(addr string, port int,user string ,password string) (*SwitchDevices, error) {
	srv := SwitchDevices{
		Host:     addr,
		Port:     port,
		ProtocolType: "tcp",
		Password:password,
		UserName:user,
	}

	return &srv, nil
}

func (dec *SwitchClient)NewSShClient(d *SwitchDevices) (*SwitchClient, error) {
	config := &SwitchClient{
		ClientConfig: &ssh.ClientConfig{
			Timeout:         time.Second * 5,
			User:            d.UserName,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Config : ssh.Config{
				Ciphers: []string{"aes128-cbc"},
			},
		},
	}

	// 设置设备访问密码，如果是keyFile ,password 为访问文件路径
	config = config.SetPassword(PasswordString,d.Password)
	addr := fmt.Sprintf("%s:%d", d.Host, d.Port)
	c, err := ssh.Dial("tcp", addr, config.ClientConfig)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	dec.Client = c
	return dec, nil

}

func (dec *SwitchClient)SetPassword(sshType int,password string) *SwitchClient {

	if sshType == PasswordString {
		dec.ClientConfig.Auth = []ssh.AuthMethod{ssh.Password(password)}
	}
	if sshType == PasswordKeyFile {
		dec.ClientConfig.Auth = []ssh.AuthMethod{publicKeyAuthFunc(password)}
	}
	return dec
}

func (dec *SwitchClient) DeviceCmd(cmd []string)  string {
	var res strings.Builder
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4k baud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4k baud
	}

	session, err := dec.Client.NewSession()
	defer session.Close()

	if err := session.RequestPty("vt100", 80, 40, modes); err != nil {
		fmt.Println(err)
	}

	w, err := session.StdinPipe()
	if err != nil {
		panic(err)
	}
	r, err := session.StdoutPipe()
	if err != nil {
		panic(err)
	}

	in, out := MuxShell(w, r)
	if err := session.Shell(); err != nil {
		fmt.Println(err)
	}

	<-out //ignore the shell output
	for _,c := range cmd{
		in <- c
	}

	//fmt.Printf("%s\n %s\n", <-out, <-out)
	for ou := range out {
		res.WriteString(ou)
		fmt.Printf("%s\n ", ou)
	}

	session.Wait()
	return res.String()

}

func MuxShell(w io.Writer, r io.Reader) (chan<- string, <-chan string) {
	in := make(chan string, 3)
	out := make(chan string, 5)
	var wg sync.WaitGroup
	wg.Add(1) //for the shell itself
	go func() {
		for cmd := range in {
			wg.Add(1)
			w.Write([]byte(cmd + "\n"))
			wg.Wait()
		}
	}()

	go func() {
		var (
			buf [65 * 1024]byte
			t   int
		)
		for {
			n, err := r.Read(buf[t:])
			if err != nil {
				fmt.Println(err.Error())
				close(in)
				close(out)
				return
			}
			t += n
			result := string(buf[:t])
			if strings.Contains(result, "Username:") ||
				strings.Contains(result, "Password:") ||
				strings.Contains(result, "#") {
				out <- string(buf[:t])
				t = 0
				wg.Done()
			}
		}
	}()
	return in, out
}


