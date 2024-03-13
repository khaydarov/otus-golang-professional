package memorystorage

import (
	"sync"
	"time"

	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	mu sync.RWMutex

	events map[storage.EventID]storage.Event
}

func New() *Storage {
	return &Storage{
		events: make(map[storage.EventID]storage.Event),
	}
}

func (s *Storage) Insert(event storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[event.ID]; ok {
		return storage.ErrEventAlreadyExists
	}

	s.events[event.ID] = event
	return nil
}

func (s *Storage) Update(event storage.Event) error {
	if _, ok := s.events[event.ID]; !ok {
		return storage.ErrEventDoesNotExist
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.events[event.ID] = event

	return nil
}

func (s *Storage) Delete(id storage.EventID) error {
	if _, ok := s.events[id]; !ok {
		return storage.ErrEventDoesNotExist
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.events, id)
	return nil
}

func (s *Storage) GetByID(id storage.EventID) (storage.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	event, ok := s.events[id]
	if !ok {
		return storage.Event{}, storage.ErrEventDoesNotExist
	}

	return event, nil
}

func (s *Storage) GetAll() []storage.Event {
	result := make([]storage.Event, 0, len(s.events))
	for _, event := range s.events {
		result = append(result, event)
	}
	return result
}

func (s *Storage) GetForTheDay(datetime time.Time) []storage.Event {
	var result []storage.Event
	for _, event := range s.events {
		if event.StartDate.Day() == datetime.Day() {
			result = append(result, event)
		}
	}

	return result
}

func (s *Storage) GetForTheWeek(datetime time.Time) []storage.Event {
	var result []storage.Event
	for _, event := range s.events {
		_, eventWeek := event.StartDate.ISOWeek()
		_, targetDayWeek := datetime.ISOWeek()

		if eventWeek == targetDayWeek {
			result = append(result, event)
		}
	}

	return result
}

func (s *Storage) GetForTheMonth(datetime time.Time) []storage.Event {
	var result []storage.Event
	for _, event := range s.events {
		if event.StartDate.Month() == datetime.Month() {
			result = append(result, event)
		}
	}

	return result
}

func (s *Storage) IsTimeBusy(datetime time.Time) bool {
	for _, event := range s.events {
		if event.StartDate.Day() == datetime.Day() && event.StartDate.Hour() == datetime.Hour() {
			return true
		}
	}

	return false
}
