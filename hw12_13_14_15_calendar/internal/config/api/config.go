package api

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	LogLevel    string     `yaml:"logLevel"`
	HTTPServer  HTTPServer `yaml:"httpServer"`
	StorageType string     `yaml:"storageType"`
}

type HTTPServer struct {
	Host string `yaml:"host" env-default:"localhost" env:"HTTP_HOST" env-description:"HTTP server host"`
	Port int    `yaml:"port" env-default:"8080" env:"HTTP_PORT" env-description:"HTTP server port"`
}

func Load(configFile string) (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(configFile, &cfg); err != nil {
		return nil, fmt.Errorf("cannot read config file: %w", err)
	}

	return &cfg, nil
}
