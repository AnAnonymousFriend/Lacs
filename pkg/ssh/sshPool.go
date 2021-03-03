package ssh

import (
	"golang.org/x/crypto/ssh"
	"sync"
	"time"

	"fmt"
)


type PoolConn struct {
	client   *ssh.Client
	mu       sync.RWMutex
	device   *Devices
	unusable bool
}

type Devices struct {
	sshHost     string
	sshUser     string
	sshPassword string
	sshPort     int
}

func CreateSshClient(d *Devices) *ssh.Client {
	config := &ssh.ClientConfig{
		Timeout:         time.Second,
		User:            d.sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	config.Auth = []ssh.AuthMethod{ssh.Password(d.sshPassword)}

	addr := fmt.Sprintf("%s:%d", d.sshHost, d.sshPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err !=nil {
		fmt.Println(err)
		return nil
	}
	return sshClient
}