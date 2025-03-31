// Package config defines project configurations
package config

import (
	"fmt"
	"os"

	"github.com/ezex-io/ezex-notification/api/grpc"
	"github.com/ezex-io/ezex-notification/internal/adapters/smtp"
	"gopkg.in/yaml.v3"
)

type Config struct {
	SMTP *smtp.Config `yaml:"smtp"`
	GRPC *grpc.Config `yaml:"grpc"`
}

func Load(path string) (*Config, error) {
	payload, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load config file: %w", err)
	}

	var config Config

	config.SMTP = smtp.DefaultConfig()
	config.GRPC = grpc.DefaultConfig()

	if err := yaml.Unmarshal(payload, config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal configurations: %w", err)
	}

	if err := config.BasicCheck(); err != nil {
		return nil, fmt.Errorf("confgurations basic check failed: %w", err)
	}

	return &config, nil
}

// BasicCheck checks the necessary config checking from each module.
func (cfg *Config) BasicCheck() error {
	if err := cfg.SMTP.BasicCheck(); err != nil {
		return err
	}

	if err := cfg.GRPC.BasicCheck(); err != nil {
		return err
	}

	return nil
}
