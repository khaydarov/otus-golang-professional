package model

import (
	"time"

	"github.com/google/uuid"
)

type EventID struct {
	id string
}

func (e *EventID) String() string {
	return e.id
}

func (e *EventID) Value() string {
	return e.id
}

func NewEventID() EventID {
	uuidv7, _ := uuid.NewV7()
	return EventID{id: uuidv7.String()}
}

func CreateEventIDFrom(id string) EventID {
	return EventID{id: id}
}

type Event struct {
	ID          EventID
	CreatorID   string
	Title       string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	NotifyAt    time.Time
}

type Events []Event
