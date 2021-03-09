package ssh

type Devices interface {
	NewDevcie()
	NewSShClient()
	DeviceCmd()
}


const (
	PasswordString = iota
	PasswordKeyFile
)
