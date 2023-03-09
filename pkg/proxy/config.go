package proxy

import "github.com/kelseyhightower/envconfig"

// Config is a struct for configuring the proxy server.
type Config struct {
	// Port is the port number for the reverse proxy.
	Port int `envconfig:"PORT" default:"8080"`
	// APIHost is the destination host to which requests will be proxied.
	APIHost string `envconfig:"API_HOST" required:"true"`
}

// GetConfig returns a Config struct populated using env variables.
func GetConfig() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
