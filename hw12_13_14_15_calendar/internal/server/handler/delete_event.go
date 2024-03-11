package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type DeleteEventRequest struct {
	ID string `json:"id"`
}

type DeleteEventResponse struct {
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
