package memorystorage

import (
	"sync"

	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/domain"
)

type Storage struct {
	// TODO
	mu sync.RWMutex //nolint:unused
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Create(event *domain.Event) (string, error) {
	return "", nil
}

func (s *Storage) Update(event *domain.Event) error {
	return nil
}

func (s *Storage) Delete(id string) error {
	return nil
}

func (s *Storage) GetAll() ([]*domain.Event, error) {
	return nil, nil
}

func (s *Storage) GetByID(id string) (*domain.Event, error) {
	return nil, nil
}
