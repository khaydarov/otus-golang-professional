package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/app"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/storage/memory"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "../../configs/config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	cfg := config.MustLoad(configFile)
	log := logger.New(cfg.Env)

	storage := memorystorage.New()
	calendar := app.New(log, storage)

	server := internalhttp.NewServer(&cfg.HTTPServer, log, calendar)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			log.Error("failed to stop http server: " + err.Error())
		}
	}()

	log.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		log.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
