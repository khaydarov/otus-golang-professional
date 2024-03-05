package app

import (
	"context"
	"log/slog"

	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/domain"
)

type App struct {
	log     *slog.Logger
	storage Storage
}

type Storage interface {
	Create(event *domain.Event) (string, error)
}

func New(logger *slog.Logger, storage Storage) *App {
	return &App{
		logger,
		storage,
	}
}

func (a *App) CreateEvent(_ context.Context, id, title string) (string, error) {
	id, err := a.storage.Create(&domain.Event{ID: id, Title: title})
	if err != nil {
		return "", err
	}

	return id, nil
}
