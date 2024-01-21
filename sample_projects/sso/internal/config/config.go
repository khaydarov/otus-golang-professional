package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"time"
)

type Config struct {
	Env  string `env:"ENV"`
	GRPC GRPCConfig
}

type GRPCConfig struct {
	Port    int           `env:"PORT"`
	Timeout time.Duration `env:"TIMEOUT" env-default:"1h"`
}

func MustLoad() *Config {
	var cfg Config

	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		log.Fatalf("cannot read env file: %s", err)
	}

	return &cfg
}
