package config

// Server config
type Server struct {
	Host string `yaml:"host" env:"HOST" env-default:"0.0.0.0`
	Port string `yaml:"port" env:"PORT"`
}

// NewServer ...
func NewServer() *Server {
	return &Server{}
}
