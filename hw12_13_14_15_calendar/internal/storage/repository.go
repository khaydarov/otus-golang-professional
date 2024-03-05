package storage

import (
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/domain"
	"time"
)

type EventRepository interface {
	Create(event *domain.Event) error
	Update(event *domain.Event) error
	Delete(id *domain.EventID) error
	GetAll() []*domain.Event
	GetForTheDay(day time.Time) []*domain.Event
	GetForTheWeek(day time.Time) []*domain.Event
	GetForTheMonth(day time.Time) []*domain.Event
}
