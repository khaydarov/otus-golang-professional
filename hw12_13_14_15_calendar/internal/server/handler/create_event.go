package handler

import (
	"encoding/json"
	"net/http"
)

type CreateEventRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	NotifyAt    string `json:"notifyAt"`
}

type CreateEventResponse struct {
	ID string `json:"id"`
}

type EventCreator interface {
	CreateEvent(title, description, creatorID, startDate, endDate, notify string) (string, error)
}

func CreateEventHandler(c EventCreator) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateEventRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := c.CreateEvent(
			req.Title,
			req.Description,
			"8063e703-7d3d-495e-a014-968ade36dc3f", // get user id from request
			req.StartDate,
			req.EndDate,
			req.NotifyAt,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(id))
	}
}
