package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/khaydarovm/otus-golang-professional/sample_projects/url-shortener/internal/config"
	mwLogger "github.com/khaydarovm/otus-golang-professional/sample_projects/url-shortener/internal/http-server/middleware"
	"github.com/khaydarovm/otus-golang-professional/sample_projects/url-shortener/internal/logger"
	"github.com/khaydarovm/otus-golang-professional/sample_projects/url-shortener/internal/logger/sl"
	"github.com/khaydarovm/otus-golang-professional/sample_projects/url-shortener/internal/storage/inmemory"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)

	storage, err := inmemory.New()
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})

	server := &http.Server{
		Addr:    fmt.Sprintf("localhost:%s", cfg.Port),
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Error("failed to start server: %s", err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	slog.Info("server started")

	<-done
	slog.Info("stopping server")
}
