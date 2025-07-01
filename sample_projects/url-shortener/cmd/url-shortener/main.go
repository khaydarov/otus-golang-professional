package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/khaydarov/otus-golang-professional/sample_projects/url-shortener/internal/clients/sso/grpc"
	"github.com/khaydarov/otus-golang-professional/sample_projects/url-shortener/internal/config"
	"github.com/khaydarov/otus-golang-professional/sample_projects/url-shortener/internal/http-server/handlers/url/save"
	mwCustom "github.com/khaydarov/otus-golang-professional/sample_projects/url-shortener/internal/http-server/middleware"
	"github.com/khaydarov/otus-golang-professional/sample_projects/url-shortener/internal/lib/logger"
	"github.com/khaydarov/otus-golang-professional/sample_projects/url-shortener/internal/storage/inmemory"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)

	ssoclient, err := grpc.New(context.Background(), log, "localhost", 10, 3)
	if err != nil {
		log.Error("failed to create sso client: %s", err)
		os.Exit(1)
	}

	_ = ssoclient

	storage := inmemory.New()

	router := chi.NewRouter()
	router.Use(mwCustom.RequestID)
	router.Use(mwCustom.New(log))
	router.Use(middleware.Recoverer)

	router.Post("/", save.New(log, storage))

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
