package config

import (
	"log"

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

func MustLoad(configFile string) *Config {
	var cfg Config

	if err := cleanenv.ReadConfig(configFile, &cfg); err != nil {
		log.Fatalf("cannot read config file: %s", err)
	}

	return &cfg
}
