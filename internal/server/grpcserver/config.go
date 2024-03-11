package grpcserver

import "fmt"

type Config struct {
	Port       string `yaml:"port"`
	CertPath   string `yaml:"cert_path"`
	KeyPath    string `yaml:"key_path"`
	CACertPath string `yaml:"ca_cert_path"`
}

func (c *Config) Adress() string {
	return fmt.Sprintf(":%s", c.Port)
}
