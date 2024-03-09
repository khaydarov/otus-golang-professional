package memorystorage

import (
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
)

func TestStorageCreate(t *testing.T) {
	s := New()
	event := generateEvents(1)
	err := s.Insert(event[0])

	require.NoError(t, err)
	require.Len(t, s.events, 1)
}

func TestStorageUpdate(t *testing.T) {
	s := New()
	event := generateEvents(1)
	err := s.Insert(event[0])

	require.NoError(t, err)
	require.Len(t, s.events, 1)

	event[0].Title = faker.Word()
	err = s.Update(event[0])

	require.NoError(t, err)
	require.Len(t, s.events, 1)
	require.Equal(t, event[0].Title, s.events[event[0].ID].Title)
}

func TestStorageDelete(t *testing.T) {
	s := New()
	event := generateEvents(1)
	err := s.Insert(event[0])

	require.NoError(t, err)
	require.Len(t, s.events, 1)

	err = s.Delete(event[0].ID)

	require.NoError(t, err)
	require.Len(t, s.events, 0)
}

func TestStorageGetAll(t *testing.T) {
	s := New()
	events := generateEvents(5)
	for _, event := range events {
		err := s.Insert(event)
		require.NoError(t, err)
	}

	allEvents := s.GetAll()
	require.Len(t, allEvents, 5)
}

func TestStorageGetAllEmpty(t *testing.T) {
	s := New()
	allEvents := s.GetAll()
	require.Len(t, allEvents, 0)
}

func TestStorageCreateEventAlreadyExists(t *testing.T) {
	s := New()
	event := generateEvents(1)
	err := s.Insert(event[0])

	require.NoError(t, err)
	require.Len(t, s.events, 1)

	err = s.Insert(event[0])

	require.Error(t, err)
	require.Equal(t, storage.ErrEventAlreadyExists, err)
}

func TestStorageGetForTheDay(t *testing.T) {
	s := New()
	events := generateEvents(10)
	for _, event := range events {
		err := s.Insert(event)
		require.NoError(t, err)
	}

	date, _ := time.Parse("2006-01-02", "2024-01-01")
	allEvents := s.GetForTheDay(date)
	for _, event := range allEvents {
		require.Equal(t, date.Day(), event.StartDate.Day())
	}
}

func TestStorageGetForTheWeek(t *testing.T) {
	s := New()
	events := generateEvents(10)
	for _, event := range events {
		s.Insert(event)
	}

	date, _ := time.Parse("2006-01-02", "2024-01-01")
	allEvents := s.GetForTheWeek(date)
	for _, event := range allEvents {
		_, dateWeek := date.ISOWeek()
		_, eventWeek := event.StartDate.ISOWeek()
		require.Equal(t, dateWeek, eventWeek)
	}
}

func generateEvents(count int) []storage.Event {
	events := make([]storage.Event, count)
	for i := 0; i < count; i++ {
		eventTime, _ := time.Parse(time.DateTime, faker.Timestamp())
		events[i] = storage.Event{
			ID:          storage.CreateEventIDFrom(faker.UUIDHyphenated()),
			Title:       faker.Word(),
			StartDate:   eventTime,
			EndDate:     eventTime.Add(time.Hour),
			Description: faker.Sentence(),
			UserID:      "10",
			Notify:      time.Hour,
		}
	}
	return events
}
