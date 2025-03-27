package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/api/event/handler"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/api/event/middleware"
)

func NewRouter(app Application) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.LoggerMiddleware("./logs/log.txt"))
	router.Route("/events", func(r chi.Router) {
		r.Post("/", handler.CreateEventHandler(app))
		r.Post("/{id}", handler.UpdateEventHandler(app))
		r.Delete("/{id}", handler.DeleteEventHandler(app))
		r.Get("/", handler.GetEventsHandler(app))
	})

	return router
}
