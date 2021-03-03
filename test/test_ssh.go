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

type Option func(*Device)

func Protocol(p string) Option {
	return func(s *Device) {
		s.Protocol = p
	}
}

type clientConfig ssh.ClientConfig

func (con clientConfig)SetPassword(sshType int,password string)  {
	if sshType == 0 {
		con.Auth = []ssh.AuthMethod{ssh.Password(password)}
	} else {
		con.Auth = []ssh.AuthMethod{publicKeyAuthFunc(password)}
	}

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
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}


	config.Auth = []ssh.AuthMethod{ssh.Password(d.Password)}
	addr := fmt.Sprintf("%s:%d", d.Host, d.Port)
	c, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		println(err)
		return nil, err
	}
	return c, nil
}