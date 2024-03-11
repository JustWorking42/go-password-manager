package session

type Config struct {
	SessionPath string `yaml:"session_path"`
	DBSecret    string `yaml:"db_secret"`
}
