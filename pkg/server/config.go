package server

import (
	"crypto/tls"
	"fmt"
)

// Config is standard configuration of most server commands
type Config struct {
	Bind   string `config:"bind,default=0.0.0.0"`
	Listen uint   `config:"listen,default=5000"`
	TLS    TLSConfig
}

func (c Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Bind, c.Listen)
}

type TLSConfig struct {
	ClientAuth tls.ClientAuthType
	CertFile   string
	KeyFile    string
	RootCAFile string
}

func (c TLSConfig) GetEnabled() bool {
	return len(c.CertFile) > 0 || len(c.KeyFile) > 0
}
