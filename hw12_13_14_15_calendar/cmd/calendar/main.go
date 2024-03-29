package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/app"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/storage/sql"
)

var configFile string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error loading .env file")
	}

	flag.StringVar(&configFile, "config", "configs/config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	cfg, err := config.Load(configFile)
	if err != nil {
		log.Fatalln("failed to load config")
	}

	logg := logger.New(cfg.LogLevel)
	storage, err := initStorage(cfg.StorageType)
	if err != nil {
		log.Fatalln("failed to init storage")
	}

	calendar := app.New(logg, storage)
	server := internalhttp.NewServer(&cfg.HTTPServer, logg, calendar)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}

func initStorage(storageType string) (app.Storage, error) {
	if storageType == "memory" {
		return memorystorage.New(), nil
	}

	storage := sqlstorage.New()
	err := storage.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return storage, nil
}
