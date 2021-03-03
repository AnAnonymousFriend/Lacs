package main

import (

	"fmt"
	"golang.org/x/crypto/ssh"
	"time"
)

type Device struct {
	sshHost     string
	sshUser     string
	sshPassword string
	sshPort     int
	sshType  string
	sshKeyPath string
}


func main()  {
	d := Device{
		sshHost: "172.168.1.110",
		sshUser: "root",
		sshPassword :"root",
		sshPort: 22,
	}
	client,err := NewSshClient(&d)
	if err !=nil {
		println("连接失败")
	}
	if client !=nil {
		println("已连接")
	}
}

func NewSshClient(d *Device) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		Timeout:         time.Second * 5,
		User:            d.sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	config.Auth = []ssh.AuthMethod{ssh.Password(d.sshPassword)}
	addr := fmt.Sprintf("%s:%d", d.sshHost, d.sshPort)
	c, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		println(err)
		return nil, err
	}
	return c, nil
}