package event

import (
	"errors"
	"time"
)

var (
	ErrEventAlreadyExists = errors.New("event with such ID already exists")
	ErrEventDoesNotExist  = errors.New("event with such ID does not exist")
)

type EventRepository interface {
	Insert(event Event) error
	Update(event Event) error
	Delete(id EventID) error
	GetByID(id EventID) (Event, error)
	GetAll() Events
	GetForTheDay(datetime time.Time) Events
	GetForTheWeek(datetime time.Time) Events
	GetForTheMonth(datetime time.Time) Events
	IsTimeBusy(datetime time.Time) bool
}
