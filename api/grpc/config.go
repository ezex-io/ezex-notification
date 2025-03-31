package grpc

import "errors"

type Config struct {
	Port string `yaml:"port"`
}

func DefaultConfig() *Config {
	return &Config{
		Port: "50051",
	}
}

func (c *Config) BasicCheck() error {
	if c.Port == "" {
		return errors.New("config: gRPC port dose not set")
	}

	return nil
}
