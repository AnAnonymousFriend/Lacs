package ssh

type Devices interface {
	NewSShClient()
	DeviceCmd()
}

const (
	PasswordString = iota
	PasswordKeyFile
)
