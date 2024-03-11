// Package storage provides a way to load configuration from files for storage.
package storage

import "fmt"

// Config is the configuration for storage.
type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Secret   string `yaml:"secret"`
}

// Uri returns the URI for the storage.
func (c *Config) Uri() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d", c.User, c.Password, c.Host, c.Port)
}
