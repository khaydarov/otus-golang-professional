package memorystorage

import (
	"errors"
	"sync"
	"time"

	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/domain"
)

var (
	ErrEventAlreadyExists = errors.New("event with such ID already exists")
	ErrEventDoesNotExist  = errors.New("event with such ID does not exist")
)

type Storage struct {
	mu sync.RWMutex //nolint:unused

	events map[*domain.EventID]*domain.Event
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Create(event *domain.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[&event.ID]; ok {
		return ErrEventAlreadyExists
	}

	s.events[&event.ID] = event

	return nil
}

func (s *Storage) Update(event *domain.Event) error {
	if _, ok := s.events[&event.ID]; !ok {
		return ErrEventDoesNotExist
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.events[&event.ID] = event

	return nil
}

func (s *Storage) Delete(id *domain.EventID) error {
	if _, ok := s.events[id]; !ok {
		return ErrEventDoesNotExist
	}

	s.mu.Lock()
	s.mu.Unlock()
	delete(s.events, id)

	return nil
}

func (s *Storage) GetAll() []*domain.Event {
	result := make([]*domain.Event, 0, len(s.events))
	for _, event := range s.events {
		result = append(result, event)
	}
	return result
}

func (s *Storage) GetForTheDay(day time.Time) []*domain.Event {
	return []*domain.Event{}
}

func (s *Storage) GetForTheWeek(day time.Time) []*domain.Event {
	return []*domain.Event{}
}

func (s *Storage) GetForTheMonth(day time.Time) []*domain.Event {
	return []*domain.Event{}
}
