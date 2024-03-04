package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	Env        string     `yaml:"env"`
	HttpServer HttpServer `yaml:"httpServer"`
}

type HttpServer struct {
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
