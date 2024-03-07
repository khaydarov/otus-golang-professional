package storage

import (
	"errors"
	"time"
)

var (
	ErrEventAlreadyExists = errors.New("event with such ID already exists")
	ErrEventDoesNotExist  = errors.New("event with such ID does not exist")
	ErrDateBusy           = errors.New("date is busy")
)

type EventRepository interface {
	Insert(event Event) error
	Update(event Event) error
	Delete(id EventID) error
	GetAll() []Event
	GetForTheDay(day time.Time) []Event
	GetForTheWeek(day time.Time) []Event
	GetForTheMonth(day time.Time) []Event
}
