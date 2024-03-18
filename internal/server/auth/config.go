// Package auth provides a way to load configuration from files for authentication.
package auth

// Config is the authentication configuration.
type Config struct {
	Secret string `yaml:"secret"`
}
