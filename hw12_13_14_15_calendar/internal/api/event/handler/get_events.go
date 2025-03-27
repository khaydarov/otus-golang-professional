package handler

import (
	"fmt"
	"net/http"

	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/model"
)

type EventsRetriever interface {
	GetEventsForTheDay(date string) model.Events
	GetEventsForTheWeek(date string) model.Events
	GetEventsForTheMonth(date string) model.Events
}

// GetEventsHandler returns a handler for getting events
// It handles requests to the following endpoints:
// - /events?filter=day&date=2021-10-10.
// - /events?filter=week&date=2021-10-10.
// - /events?filter=month&date=2021-10-10.
func GetEventsHandler(e EventsRetriever) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		filter := r.URL.Query().Get("filter")
		date := r.URL.Query().Get("date")

		var events model.Events

		switch filter {
		case "day":
			events = e.GetEventsForTheDay(date)
		case "week":
			events = e.GetEventsForTheWeek(date)
		case "month":
			events = e.GetEventsForTheMonth(date)
		default:
			http.Error(w, "invalid filter", http.StatusBadRequest)
			return
		}

		fmt.Println(events)
		// write response
		w.WriteHeader(http.StatusOK)
	}
}
