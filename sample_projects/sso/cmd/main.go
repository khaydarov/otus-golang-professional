package main

import (
	"fmt"
	"github.com/khaydarovm/otus-golang-professional/sample_projects/sso/internal/config"
	"github.com/khaydarovm/otus-golang-professional/sample_projects/sso/internal/lib/logger"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)

	_ = log

	fmt.Println(cfg)
}
