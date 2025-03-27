package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UpdateEventRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Notify      string `json:"notify"`
}

type UpdateEventResponse struct {
	ID string `json:"id"`
}

type EventUpdater interface {
	UpdateEvent(id, title, description, startDate, endDate, notify string) error
}

func UpdateEventHandler(u EventUpdater) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UpdateEventRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "failed to decode request: "+err.Error(), http.StatusBadRequest)
			return
		}

		eventID := chi.URLParam(r, "id")
		err := u.UpdateEvent(
			eventID,
			req.Title,
			req.Description,
			req.StartDate,
			req.EndDate,
			req.Notify,
		)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(eventID))
	}
}
