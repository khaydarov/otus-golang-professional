package app

import (
	grpcapp "github.com/khaydarov/otus-golang-professional/sample_projects/sso/internal/app/grpc"
	"github.com/khaydarov/otus-golang-professional/sample_projects/sso/internal/services/auth"
	"github.com/khaydarov/otus-golang-professional/sample_projects/sso/internal/storage/inmemory"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, port int, tokenTTL time.Duration) *App {
	storage := inmemory.New()

	authService := auth.New(log, storage)
	grpcApp := grpcapp.New(
		log,
		port,
		authService,
	)

	return &App{
		grpcApp,
	}
}
