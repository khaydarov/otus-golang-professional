package main

import (
	"github.com/khaydarov/otus-golang-professional/sample_projects/sso/internal/config"
	"github.com/khaydarov/otus-golang-professional/sample_projects/sso/internal/lib/logger"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)

	_ = log
}
