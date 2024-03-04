package storage

import "github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/domain"

type EventRepository interface {
	Create(event *domain.Event) (string, error)
	Update(event *domain.Event) error
	Delete(id string) error
	GetAll() ([]*domain.Event, error)
	GetByID(id string) (*domain.Event, error)
}
