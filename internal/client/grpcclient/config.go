package grpcclient

type Config struct {
	Address    string `yaml:"address"`
	CertPath   string `yaml:"cert_path"`
	KeyPath    string `yaml:"key_path"`
	CACertPath string `yaml:"ca_cert_path"`
}
