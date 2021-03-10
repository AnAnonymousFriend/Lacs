package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

// 这里的交换机设备代指 所有网络设备，例如：网关,AC,AP,交换机等

type SwitchClient struct {
	Client    *ssh.Client
	ClientConfig *ssh.ClientConfig
	Devices *SwitchDevices
	unusable bool
}

const (
	PasswordStringa = iota
	PasswordKeyFilea
)

type SwitchDevices struct {
	Host     string // IP 地址
	Port     int	   // 端口 22
	UserName     string // 用户名
	Password string // 密码
	ProtocolType string //协议类型
}

func main()  {
	oneDevice,_ := newSwitchDevices("172.168.1.24",22,"admin","fs.com123")
	oneClient := &SwitchClient{
		Devices: oneDevice,
	}
	oneClient,_ =  oneClient.NewSShClient(oneDevice)
	if oneClient != nil {
		oneClient.DeviceCmd()
	}

}

func newSwitchDevices(addr string, port int,user string ,password string) (*SwitchDevices, error) {
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
	config = config.SetPassword(PasswordStringa,d.Password)
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

	if sshType == PasswordStringa {
		dec.ClientConfig.Auth = []ssh.AuthMethod{ssh.Password(password)}
	}
	if sshType == PasswordKeyFilea {
		dec.ClientConfig.Auth = []ssh.AuthMethod{newpublicKeyAuthFunc(password)}
	}
	return dec
}

func (dec *SwitchClient) DeviceCmd()  {
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echoingPasswordString
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
	e, err := session.StderrPipe()
	if err != nil {
		panic(err)
	}

	in, out := muxShell(w, r, e)
	if err := session.Shell(); err != nil {
		fmt.Println(err)
	}
	<-out //ignore the shell output
	in <- "show arp"
	in <- "show int status"

	in <- "exit"
	in <- "exit"
	fmt.Printf("%s\n%s\n", <-out, <-out)
	_, _ = <-out, <-out
	session.Wait()
}

func muxShell(w io.Writer, r, e io.Reader) (chan<- string, <-chan string) {
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

// Get KeyFile Password
func newpublicKeyAuthFunc(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	if err != nil {
		println("find key's home dir failed", err)
	}
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		println("ssh key file read failed", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		println("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}


