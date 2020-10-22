package config

// Server config
type Server struct {
	Host string `yaml:"host" env:"HOST" env-default:"0.0.0.0` // nolint
	Port string `yaml:"port" env:"PORT"`

	JokeURL string `yaml:"joke-url" env:"JOKE_URL"`
}

// NewServer ...
func NewServer() *Server {
	return &Server{}
}
