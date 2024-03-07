package memorystorage

import (
	"fmt"
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

	// check if date is busy
	for _, e := range s.events {
		if e.DateTime.Day() == event.DateTime.Day() && e.DateTime.Hour() == event.DateTime.Hour() {
			fmt.Println(e.DateTime, event.DateTime)
			return storage.ErrDateBusy
		}
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

func (s *Storage) GetAll() []storage.Event {
	result := make([]storage.Event, 0, len(s.events))
	for _, event := range s.events {
		result = append(result, event)
	}
	return result
}

func (s *Storage) GetForTheDay(day time.Time) []storage.Event {
	var result []storage.Event
	for _, event := range s.events {
		if event.DateTime.Day() == day.Day() {
			result = append(result, event)
		}
	}

	return result
}

func (s *Storage) GetForTheWeek(day time.Time) []storage.Event {
	var result []storage.Event
	for _, event := range s.events {
		_, eventWeek := event.DateTime.ISOWeek()
		_, targetDayWeek := day.ISOWeek()

		if eventWeek == targetDayWeek {
			result = append(result, event)
		}
	}

	return result
}

func (s *Storage) GetForTheMonth(day time.Time) []storage.Event {
	var result []storage.Event
	for _, event := range s.events {
		if event.DateTime.Month() == day.Month() {
			result = append(result, event)
		}
	}

	return result
}
