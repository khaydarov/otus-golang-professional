package internalhttp

import (
	"context"
	"fmt"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/server/handler"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
)

type Server struct {
	cfg *config.HTTPServer
	app Application
}

type Application interface {
	CreateEvent(title, description, creatorID, startDate, endDate, notify string) (string, error)
	DeleteEvent(id string) error
	UpdateEvent(id, title, description, startDate, endDate, notify string) error
}

func NewServer(httpServerCfg *config.HTTPServer, app Application) *Server {
	return &Server{
		httpServerCfg,
		app,
	}
}

func (s *Server) Start(ctx context.Context) error {
	router := chi.NewRouter()

	router.Use(LoggerMiddleware("./logs/log.txt"))
	router.Route("/events", func(r chi.Router) {
		r.Post("/", handler.CreateEventHandler(s.app))
		r.Post("/{id}", handler.UpdateEventHandler(s.app))
		r.Delete("/{id}", handler.DeleteEventHandler(s.app))
	})

	//router.Post("/events/{id}", handler.UpdateEventHandler(s.app))
	//router.Get("/events", handler.GetEventsHandler(s.app))

	server := &http.Server{
		Addr:        fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port),
		ReadTimeout: 0,
		Handler:     router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("failed to start server: %s", err)
		}
	}()

	<-ctx.Done()
	return nil
}

func (s *Server) Stop(_ context.Context) error {
	return nil
}
