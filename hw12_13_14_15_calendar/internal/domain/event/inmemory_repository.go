package event

import (
	"sync"
	"time"
)

type InMemoryRepository struct {
	mu sync.RWMutex

	events map[EventID]Event
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		events: make(map[EventID]Event),
	}
}

func (s *InMemoryRepository) Insert(event Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.events[event.ID]; ok {
		return ErrEventAlreadyExists
	}

	s.events[event.ID] = event
	return nil
}

func (s *InMemoryRepository) Update(event Event) error {
	if _, ok := s.events[event.ID]; !ok {
		return ErrEventDoesNotExist
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.events[event.ID] = event

	return nil
}

func (s *InMemoryRepository) Delete(id EventID) error {
	if _, ok := s.events[id]; !ok {
		return ErrEventDoesNotExist
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.events, id)
	return nil
}

func (s *InMemoryRepository) GetByID(id EventID) (Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	event, ok := s.events[id]
	if !ok {
		return Event{}, ErrEventDoesNotExist
	}

	return event, nil
}

func (s *InMemoryRepository) GetAll() Events {
	result := make([]Event, 0, len(s.events))
	for _, event := range s.events {
		result = append(result, event)
	}
	return result
}

func (s *InMemoryRepository) GetForTheDay(datetime time.Time) Events {
	var result Events
	for _, event := range s.events {
		if event.StartDate.Day() == datetime.Day() {
			result = append(result, event)
		}
	}

	return result
}

func (s *InMemoryRepository) GetForTheWeek(datetime time.Time) Events {
	var result Events
	for _, event := range s.events {
		_, eventWeek := event.StartDate.ISOWeek()
		_, targetDayWeek := datetime.ISOWeek()

		if eventWeek == targetDayWeek {
			result = append(result, event)
		}
	}

	return result
}

func (s *InMemoryRepository) GetForTheMonth(datetime time.Time) Events {
	var result Events
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
