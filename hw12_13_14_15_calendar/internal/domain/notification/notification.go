package model

import "time"

type Notification struct {
	EventID    EventID
	EventTitle string
	EventDate  time.Time
	ReceiverID string
}

type Notifications []Notification
