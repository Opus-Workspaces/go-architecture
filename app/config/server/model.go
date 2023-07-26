package config

type ServerType struct {
	Address        string
	Port           string
	Timeout        uint16
	AppVersion     string
	ReadTimeout    uint16
	WriteTimeout   uint16
	DefaultTimeout uint16
}
