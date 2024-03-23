// Package session provides a way to load configuration from files.
package session

// Config is the configuration for a session manager.
type Config struct {
	SessionPath string `yaml:"session_path"`
	DBSecret    string `yaml:"db_secret"`
}
