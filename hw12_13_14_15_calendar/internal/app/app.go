package app

import (
	"context"
	"log/slog"

	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/storage"
)

type App struct {
	log *slog.Logger
	r   Storage
}

type Storage interface {
	Insert(event storage.Event) error
}

func New(logger *slog.Logger, s Storage) *App {
	return &App{
		logger,
		s,
	}
}

func (a *App) CreateEvent(_ context.Context, title string) (string, error) {
	newEvent := storage.Event{
		ID:    storage.NewEventID(),
		Title: title,
	}

	err := a.r.Insert(newEvent)
	if err != nil {
		return "", err
	}

	return newEvent.ID.Value(), nil
}
