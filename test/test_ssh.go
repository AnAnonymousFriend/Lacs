package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"time"
)

type Device struct {
	Host     string // IP 地址
	Port     int	   // 端口 22
	UserName     string // 用户命
	Password string // 密码
	conf *DeviceConfig
	Protocol string

}

type DeviceConfig struct {
	sshType  int
	sshKeyPath string
}
const (
	PasswordString = iota    // 开始生成枚举值, 默认为0
	PasswordFile

)

func SetPassword(config *ssh.ClientConfig,sshType int,password string) *ssh.ClientConfig {
	if sshType == PasswordString {
		config.Auth = []ssh.AuthMethod{ssh.Password(password)}
	}
	if sshType == PasswordFile  {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(password)}
	}
	return config
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


func main()  {


}

func NewSshClient(d *Device) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		Timeout:         time.Second * 5,
		User:            d.UserName,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	config = SetPassword(config,0,"root")
	addr := fmt.Sprintf("%s:%d", d.Host, d.Port)
	c, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		println(err)
		return nil, err
	}
	return c, nil

}