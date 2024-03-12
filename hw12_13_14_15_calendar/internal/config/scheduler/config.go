package api

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	LogLevel string `yaml:"logLevel"`
}

func Load(configFile string) (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(configFile, &cfg); err != nil {
		return nil, fmt.Errorf("cannot read config file: %w", err)
	}

	return &cfg, nil
}
