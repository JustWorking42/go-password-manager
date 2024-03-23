// Package storage provides a way to load configuration from files for storage.
package storage

// Config is the configuration for storage.
type Config struct {
	Secret string `yaml:"secret"`
	Set    string `yaml:"set"`
}
