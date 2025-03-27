package event_repository

import (
	"sync"
	"time"

	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/model"
)

type InMemoryRepository struct {
	mu sync.RWMutex

	events map[model.EventID]model.Event
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		events: make(map[model.EventID]model.Event),
	}
}

func (s *InMemoryRepository) Insert(event model.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[event.ID]; ok {
		return ErrEventAlreadyExists
	}

	s.events[event.ID] = event
	return nil
}

func (s *InMemoryRepository) Update(event model.Event) error {
	if _, ok := s.events[event.ID]; !ok {
		return ErrEventDoesNotExist
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.events[event.ID] = event

	return nil
}

func (s *InMemoryRepository) Delete(id model.EventID) error {
	if _, ok := s.events[id]; !ok {
		return ErrEventDoesNotExist
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.events, id)
	return nil
}

func (s *InMemoryRepository) GetByID(id model.EventID) (model.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	event, ok := s.events[id]
	if !ok {
		return model.Event{}, ErrEventDoesNotExist
	}

	return event, nil
}

func (s *InMemoryRepository) GetAll() []model.Event {
	result := make([]model.Event, 0, len(s.events))
	for _, event := range s.events {
		result = append(result, event)
	}
	return result
}

func (s *InMemoryRepository) GetForTheDay(datetime time.Time) model.Events {
	var result model.Events
	for _, event := range s.events {
		if event.StartDate.Day() == datetime.Day() {
			result = append(result, event)
		}
	}

	return result
}

func (s *InMemoryRepository) GetForTheWeek(datetime time.Time) model.Events {
	var result model.Events
	for _, event := range s.events {
		_, eventWeek := event.StartDate.ISOWeek()
		_, targetDayWeek := datetime.ISOWeek()

		if eventWeek == targetDayWeek {
			result = append(result, event)
		}
	}

	return result
}

func (s *InMemoryRepository) GetForTheMonth(datetime time.Time) model.Events {
	var result model.Events
	for _, event := range s.events {
		if event.StartDate.Month() == datetime.Month() {
			result = append(result, event)
		}
	}

	return result
}

func (s *InMemoryRepository) IsTimeBusy(datetime time.Time) bool {
	for _, event := range s.events {
		if event.StartDate.Day() == datetime.Day() && event.StartDate.Hour() == datetime.Hour() {
			return true
		}
	}

	return false
}
