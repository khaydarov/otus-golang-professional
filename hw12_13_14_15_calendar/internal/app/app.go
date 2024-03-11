package app

import (
	"errors"
	"log/slog"
	"time"

	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/storage"
)

type App struct {
	log     *slog.Logger
	storage Storage
}

type Storage interface {
	Insert(event storage.Event) error
	IsTimeBusy(datetime time.Time) bool
	Delete(id storage.EventID) error
	Update(event storage.Event) error
	GetById(id storage.EventID) (storage.Event, error)
}

func New(logLevel string, storage Storage) *App {
	log := logger.New(logLevel)
	return &App{
		log,
		storage,
	}
}

func (a *App) UpdateEvent(id, title, description, startDate, endDate, notify string) error {
	event, err := a.storage.GetById(storage.CreateEventIDFrom(id))
	if err != nil {
		return err
	}

	event.Title = title
	event.Description = description
	event.StartDate, err = parseStringToTime(startDate)
	event.EndDate, err = parseStringToTime(endDate)
	event.NotifyAt, err = parseStringToTime(notify)

	err = a.storage.Update(event)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) DeleteEvent(id string) error {
	return a.storage.Delete(storage.CreateEventIDFrom(id))
}

func (a *App) CreateEvent(title, description, creatorID, startDate, endDate, notify string) (string, error) {
	startTm, err := parseStringToTime(startDate)
	if err != nil {
		return "", err
	}

	if a.storage.IsTimeBusy(startTm) {
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

	notifyAt := startTm.Add(-n)
	newEvent := storage.Event{
		ID:          storage.NewEventID(),
		CreatorID:   creatorID,
		Title:       title,
		StartDate:   startTm,
		EndDate:     endTm,
		Description: description,
		NotifyAt:    notifyAt,
	}

	err = a.storage.Insert(newEvent)
	if err != nil {
		return "", err
	}

	return newEvent.ID.Value(), nil
}

func parseStringToTime(s string) (time.Time, error) {
	t, err := time.Parse(time.DateTime, s)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
