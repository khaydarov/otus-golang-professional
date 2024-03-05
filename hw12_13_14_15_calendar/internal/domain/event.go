package domain

import (
	"github.com/google/uuid"
	"time"
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

type Event struct {
	ID          EventID
	Title       string
	Date        time.Time
	Duration    time.Duration
	Description string
	UserID      string
	Notify      time.Duration
}
