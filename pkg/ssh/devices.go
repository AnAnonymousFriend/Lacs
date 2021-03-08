package ssh

import (
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"sync"
	"time"

	"fmt"
)

const (
	PasswordString = iota
	PasswordKeyFile

)

// create New Devices
func NewDevcie(addr string, port int,user string ,password string, options ...func(*Device)) (*Device, error) {
	srv := Device{
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
func (dec *DeviceClient)SetPassword(sshType int,password string) *DeviceClient {

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
func (dec *DeviceClient)NewSShClient(d *Device) (*DeviceClient, error) {
	config := &DeviceClient{
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
func (dec *DeviceClient) DeviceCmd(shell string) (string, error) {
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


