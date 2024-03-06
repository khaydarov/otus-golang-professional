package app

import (
	"context"
	"log/slog"

	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/storage"
)

type App struct {
	log     *slog.Logger
	storage Storage
}

type Storage interface {
	Create(event storage.Event) error
}

func New(logger *slog.Logger, storage Storage) *App {
	return &App{
		logger,
		storage,
	}
}

func (a *App) CreateEvent(_ context.Context, title string) (string, error) {
	newEvent := storage.Event{
		ID:    storage.NewEventID(),
		Title: title,
	}

	err := a.storage.Create(newEvent)
	if err != nil {
		return "", err
	}

	return newEvent.ID.Value(), nil
}
