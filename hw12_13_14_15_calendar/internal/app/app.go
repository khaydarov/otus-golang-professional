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
	startDate string,
	endDate string,
	description string,
	userID string,
	notify string,
) (string, error) {
	i, err := strconv.ParseInt(startDate, 10, 64)
	if err != nil {
		return "", err
	}
	startTm := time.Unix(i, 0)

	i, err = strconv.ParseInt(endDate, 10, 64)
	if err != nil {
		return "", err
	}

	endTm := time.Unix(i, 0)
	n, err := time.ParseDuration(notify)
	if err != nil {
		return "", err
	}
	newEvent := storage.Event{
		ID:          storage.NewEventID(),
		Title:       title,
		StartDate:   startTm,
		EndDate:     endTm,
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
