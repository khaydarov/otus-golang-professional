package main

import (
	"errors"
	"log/slog"
	"time"

	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/model"
)

type Storage interface {
	Insert(event model.Event) error
	IsTimeBusy(datetime time.Time) bool
	Delete(id model.EventID) error
	Update(event model.Event) error
	GetByID(id model.EventID) (model.Event, error)
	GetForTheDay(datetime time.Time) model.Events
	GetForTheWeek(datetime time.Time) model.Events
	GetForTheMonth(datetime time.Time) model.Events
}

type Calendar struct {
	log     *slog.Logger
	storage Storage
}

func NewCalendar(storage Storage, logger *slog.Logger) *Calendar {
	return &Calendar{
		logger,
		storage,
	}
}

func (a *Calendar) GetEventsForTheDay(date string) model.Events {
	tm, err := parseStringToDateOnly(date)
	if err != nil {
		return model.Events{}
	}

	return a.storage.GetForTheDay(tm)
}

func (a *Calendar) GetEventsForTheWeek(date string) model.Events {
	tm, err := parseStringToDateOnly(date)
	if err != nil {
		return model.Events{}
	}

	return a.storage.GetForTheWeek(tm)
}

func (a *Calendar) GetEventsForTheMonth(date string) model.Events {
	tm, err := parseStringToDateOnly(date)
	if err != nil {
		return model.Events{}
	}

	return a.storage.GetForTheMonth(tm)
}

func (a *Calendar) UpdateEvent(id, title, description, startDate, endDate, notify string) error {
	event, err := a.storage.GetByID(model.CreateEventIDFrom(id))
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

func (a *Calendar) DeleteEvent(id string) error {
	return a.storage.Delete(model.CreateEventIDFrom(id))
}

func (a *Calendar) CreateEvent(title, description, creatorID, startDate, endDate, notify string) (string, error) {
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
	newEvent := model.Event{
		ID:          model.NewEventID(),
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
