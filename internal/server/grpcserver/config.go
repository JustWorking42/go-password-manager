// Package grpcserver provides a way to load configuration from files for grpc server.
package grpcserver

// Config is the grpc server configuration.
type Config struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	CertPath   string `yaml:"cert_path"`
	KeyPath    string `yaml:"key_path"`
	CACertPath string `yaml:"ca_cert_path"`
}
