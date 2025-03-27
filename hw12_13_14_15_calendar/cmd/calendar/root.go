package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/cmd/calendar/app"
	eventApi "github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/api/event"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	eventRepository "github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/repository/event"
)

var configFile string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error loading .env file")
	}

	flag.StringVar(&configFile, "config", "configs/calendar_config.yaml", "Path to configuration file")
	flag.Parse()

	logger := initLogger(os.Getenv("LOG_LEVEL"))
	cfg, err := config.Load(configFile)
	if err != nil {
		logger.Error("load config: ", err)
	}

	storage, err := initStorage(cfg.StorageType)
	if err != nil {
		logger.Error("init storage: ", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	calendar := app.NewCalendar(storage, logger)
	server := eventApi.NewServer(&cfg.HTTPServer, calendar, logger)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()

		logger.Info("calendar is running...")
		if err := server.Start(ctx); err != nil {
			logger.Debug("failed to start http server: ", err)
			cancel()
		}
	}()

	go func() {
		defer wg.Done()
		if err := server.Stop(ctx); err != nil {
			logger.Debug("failed to stop http server: ", err)
		}
	}()

	wg.Wait()
	logger.Warn("calendar is stopped")
}

func initStorage(storageType string) (app.Storage, error) {
	if storageType == "memory" {
		return eventRepository.NewInMemoryRepository(), nil
	}

	s := eventRepository.NewPsqlRepository()
	err := s.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return s, nil
}

func initLogger(logLevel string) *slog.Logger {
	return logger.New(logLevel)
}
