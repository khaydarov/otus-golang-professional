package storage

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
	GetAll() []Event
	GetForTheDay(datetime time.Time) []Event
	GetForTheWeek(datetime time.Time) []Event
	GetForTheMonth(datetime time.Time) []Event
	IsTimeBusy(datetime time.Time) bool
}
