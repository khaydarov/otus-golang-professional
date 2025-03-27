package event_repository

import (
	"errors"
	"time"

	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/model"
)

var (
	ErrEventAlreadyExists = errors.New("event with such ID already exists")
	ErrEventDoesNotExist  = errors.New("event with such ID does not exist")
)

type EventRepository interface {
	Insert(event model.Event) error
	Update(event model.Event) error
	Delete(id model.EventID) error
	GetByID(id model.EventID) (model.Event, error)
	GetAll() model.Events
	GetForTheDay(datetime time.Time) model.Events
	GetForTheWeek(datetime time.Time) model.Events
	GetForTheMonth(datetime time.Time) model.Events
	IsTimeBusy(datetime time.Time) bool
}
