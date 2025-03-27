package api_handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type DeleteEventRequest struct {
	ID string `json:"id"`
}

type EventDeleter interface {
	DeleteEvent(id string) error
}

func DeleteEventHandler(d EventDeleter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		eventID := chi.URLParam(r, "id")
		err := d.DeleteEvent(eventID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusOK)
	}
}
