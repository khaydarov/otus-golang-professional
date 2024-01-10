package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	Env  string `env:"ENV"`
	Port string `env:"PORT"`
}

func MustLoad() *Config {
	var cfg Config

	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		log.Fatalf("cannot read env file: %s", err)
	}

	return &cfg
}
