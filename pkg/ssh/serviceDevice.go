package ssh

import (
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"sync"
	"time"

	"fmt"
)

type ServiceClient struct {
	Client    *ssh.Client
	ClientConfig *ssh.ClientConfig
	Devices ServiceDevice
	unusable bool
}

type ServiceDevice struct {
	Host     string // IP 地址
	Port     int	   // 端口 22
	UserName     string // 用户名
	Password string // 密码
	ProtocolType string //协议类型
}


// create New Devices
func NewServiceDevice(addr string, port int,user string ,password string, options ...func(*ServiceDevice)) (*ServiceDevice, error) {
	srv := ServiceDevice{
		Host:     addr,
		Port:     port,
		ProtocolType: "tcp",
		Password:password,
		UserName:user,
	}
	for _, option := range options {
		option(&srv)
	}
	return &srv, nil
}

// set Devices visit to info Password
func (dec *ServiceClient)SetPassword(sshType int,password string) *ServiceClient {

	if sshType == PasswordString {
		dec.ClientConfig.Auth = []ssh.AuthMethod{ssh.Password(password)}
	}
	if sshType == PasswordKeyFile {
		dec.ClientConfig.Auth = []ssh.AuthMethod{publicKeyAuthFunc(password)}
	}
	return dec
}

// Get KeyFile Password
func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
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

// Create New SSHClient
func (dec *ServiceClient)NewSShClient(d *ServiceDevice) (*ServiceClient, error) {
	config := &ServiceClient{
		ClientConfig: &ssh.ClientConfig{
			Timeout:         time.Second * 5,
			User:            d.UserName,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
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

//  executive command
func (dec *ServiceClient) DeviceCmd(shell string) (string, error) {
	if dec.Client == nil {
		return "", nil
	}

	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	session, err := dec.Client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	buf, err := session.CombinedOutput(shell)
	return  string(buf),err
}





