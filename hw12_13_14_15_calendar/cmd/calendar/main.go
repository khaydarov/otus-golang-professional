package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/app"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	internalhttp "github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/storage/sql"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
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

	s, err := initStorage(cfg.StorageType)
	if err != nil {
		log.Fatalln("failed to init storage")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	calendar := app.New(cfg.LogLevel, s)
	server := internalhttp.NewServer(&cfg.HTTPServer, calendar)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()

		log.Println("calendar is running...")
		if err := server.Start(ctx); err != nil {
			log.Println("failed to start http server: " + err.Error())
			cancel()
		}
	}()

	go func() {
		defer wg.Done()
		if err := server.Stop(ctx); err != nil {
			log.Println("failed to stop http server: " + err.Error())
		}
	}()

	wg.Wait()
	log.Println("calendar is stopped")
}

func initStorage(storageType string) (app.Storage, error) {
	if storageType == "memory" {
		return memorystorage.New(), nil
	}

	s := sqlstorage.New()
	err := s.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return s, nil
}
