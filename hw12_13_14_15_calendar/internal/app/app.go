package app

import (
	"context"
	"errors"
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
	IsTimeBusy(start time.Time) bool
}

func New(logger *slog.Logger, r Storage) *App {
	return &App{
		logger,
		r,
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
	startTm, err := parseStringToTime(startDate)
	if err != nil {
		return "", err
	}

	if a.r.IsTimeBusy(startTm) {
		return "", errors.New("time is busy")
	}

	endTm, err := parseStringToTime(endDate)
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

func parseStringToTime(s string) (time.Time, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(i, 0), nil
}
