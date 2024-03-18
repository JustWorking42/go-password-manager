// Package grpcserver provides a way to load configuration from files for grpc server.
package grpcserver

import "fmt"

// Config is the grpc server configuration.
type Config struct {
	Port       string `yaml:"port"`
	CertPath   string `yaml:"cert_path"`
	KeyPath    string `yaml:"key_path"`
	CACertPath string `yaml:"ca_cert_path"`
}

// Adress returns the grpc server address.
func (c *Config) Adress() string {
	return fmt.Sprintf(":%s", c.Port)
}
