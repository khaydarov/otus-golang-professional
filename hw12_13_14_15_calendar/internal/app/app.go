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
	GetByID(id storage.EventID) (storage.Event, error)
	GetForTheDay(datetime time.Time) []storage.Event
	GetForTheWeek(datetime time.Time) []storage.Event
	GetForTheMonth(datetime time.Time) []storage.Event
}

func New(logLevel string, storage Storage) *App {
	log := logger.New(logLevel)
	return &App{
		log,
		storage,
	}
}

func (a *App) GetEventsForTheDay(date string) []storage.Event {
	tm, err := parseStringToDateOnly(date)
	if err != nil {
		return []storage.Event{}
	}

	return a.storage.GetForTheDay(tm)
}

func (a *App) GetEventsForTheWeek(date string) []storage.Event {
	tm, err := parseStringToDateOnly(date)
	if err != nil {
		return []storage.Event{}
	}

	return a.storage.GetForTheWeek(tm)
}

func (a *App) GetEventsForTheMonth(date string) []storage.Event {
	tm, err := parseStringToDateOnly(date)
	if err != nil {
		return []storage.Event{}
	}

	return a.storage.GetForTheMonth(tm)
}

func (a *App) UpdateEvent(id, title, description, startDate, endDate, notify string) error {
	event, err := a.storage.GetByID(storage.CreateEventIDFrom(id))
	if err != nil {
		return err
	}

	event.Title = title
	event.Description = description
	event.StartDate, _ = parseStringToTime(startDate)
	event.EndDate, _ = parseStringToTime(endDate)
	event.NotifyAt, _ = parseStringToTime(notify)

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

func parseStringToDateOnly(s string) (time.Time, error) {
	t, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
