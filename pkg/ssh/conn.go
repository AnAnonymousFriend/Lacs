package ssh

import (
	"golang.org/x/crypto/ssh"
	"sync"
)

type PoolConn struct {
	deviceName string
	mu       sync.RWMutex
	c        *DeviceClient
	unusable bool
}

type DeviceClient struct {
	Client    *ssh.Client
	ClientConfig *ssh.ClientConfig
	Devices Device
}

type Device struct {
	Host     string // IP 地址
	Port     int	   // 端口 22
	UserName     string // 用户名
	Password string // 密码
	ProtocolType string
}




