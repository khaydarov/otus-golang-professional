package main

import (
	"github.com/khaydarov/otus-golang-professional/sample_projects/sso/internal/app"
	"github.com/khaydarov/otus-golang-professional/sample_projects/sso/internal/config"
	"github.com/khaydarov/otus-golang-professional/sample_projects/sso/internal/lib/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)

	application := app.New(log, cfg.GRPC.Port, cfg.GRPC.Timeout)
	go application.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
	log.Info("application stopped")
}
