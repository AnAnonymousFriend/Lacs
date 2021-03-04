package util

import (
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"time"

	"fmt"
)


type Device struct {
	Host     string // IP 地址
	Port     int	   // 端口 22
	UserName     string // 用户命
	Password string // 密码
	conf *DeviceConfig
	Protocol string

}

const (
	PasswordString = iota    // 开始生成枚举值, 默认为0
	PasswordFile

)

type DeviceConfig struct {
	Clientconfig *ssh.ClientConfig
	Client *ssh.Client
}

func NewSshDevcie(addr string, port int,user string ,password string, options ...func(*Device)) (*Device, error) {
	srv := Device{
		Host:     addr,
		Port:     port,
		Protocol: "tcp",
		Password:password,
		UserName:user,
	}
	for _, option := range options {
		option(&srv)
	}
	return &srv, nil
}

func (d *DeviceConfig)SetPassword(sshType int,password string) *DeviceConfig {

	if sshType == PasswordString {
		d.Clientconfig.Auth = []ssh.AuthMethod{ssh.Password(password)}
	}
	if sshType == PasswordFile  {
		d.Clientconfig.Auth = []ssh.AuthMethod{publicKeyAuthFunc(password)}
	}
	return d
}

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

func NewSshClient(d *Device) (*DeviceConfig, error) {
	config := &DeviceConfig{
		Clientconfig: &ssh.ClientConfig{
			Timeout:         time.Second * 5,
			User:            d.UserName,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		},
	}

	config = config.SetPassword(0,"root")
	addr := fmt.Sprintf("%s:%d", d.Host, d.Port)
	c, err := ssh.Dial("tcp", addr, config.Clientconfig)
	if err != nil {
		println(err)
		config.Client = nil
		return nil, err

	}
	config.Client = c
	return config, nil

}


func (d DeviceConfig) SshSessionCmd(shell string) (string, error) {
	if d.Client == nil {
		return "", nil
	}
	session, err := d.Client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	buf, err := session.CombinedOutput(shell)
	return  string(buf),err
}


