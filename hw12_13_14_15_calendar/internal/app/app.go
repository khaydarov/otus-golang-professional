package app

import (
	"context"
	"log/slog"
	"strconv"
	"time"

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

func (a *App) CreateEvent(
	_ context.Context,
	title string,
	datetime string,
	duration string,
	description string,
	userID string,
	notify string,
) (string, error) {
	i, err := strconv.ParseInt(datetime, 10, 64)
	if err != nil {
		return "", err
	}

	// @todo: store unixtime instead of time
	tm := time.Unix(i, 0)
	d, err := time.ParseDuration(duration)
	if err != nil {
		return "", err
	}

	n, err := time.ParseDuration(notify)
	if err != nil {
		return "", err
	}
	newEvent := storage.Event{
		ID:          storage.NewEventID(),
		Title:       title,
		DateTime:    tm,
		Duration:    d,
		Description: description,
		UserID:      userID,
		Notify:      n,
	}

	err = a.r.Insert(newEvent)
	if err != nil {
		return "", err
	}

	return newEvent.ID.Value(), nil
}
