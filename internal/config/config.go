package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type HTTPServer struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type Config struct {
	Env         string `yaml:"env"`
	StoragePath string `yaml:"storage_path"`
	HTTPServer  `yaml:"http_server"`
}

func NewConfigFromFile(path string) (*Config, error) {
	const op = "config.NewConfigFromFile"

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("%s:\n %w", op, err)
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("%s:\n %w", op, err)
	}

	return config, nil
}
