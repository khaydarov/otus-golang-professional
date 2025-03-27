package api

import (
	"github.com/go-chi/chi/v5"
	apiHandler "github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/api/event/handler"
	apiMiddleware "github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/api/event/middleware"
)

func NewRouter(app Application) *chi.Mux {
	router := chi.NewRouter()

	router.Use(apiMiddleware.LoggerMiddleware("./logs/log.txt"))
	router.Route("/events", func(r chi.Router) {
		r.Post("/", apiHandler.CreateEventHandler(app))
		r.Post("/{id}", apiHandler.UpdateEventHandler(app))
		r.Delete("/{id}", apiHandler.DeleteEventHandler(app))
		r.Get("/", apiHandler.GetEventsHandler(app))
	})

	return router
}
