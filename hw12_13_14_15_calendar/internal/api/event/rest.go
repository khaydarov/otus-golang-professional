package api

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	eventDomain "github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/domain/event"
)

type Server struct {
	cfg    *config.HTTPServer
	app    Application
	logger *slog.Logger
}

type Application interface {
	CreateEvent(title, description, creatorID, startDate, endDate, notify string) (string, error)
	DeleteEvent(id string) error
	UpdateEvent(id, title, description, startDate, endDate, notify string) error
	GetEventsForTheDay(date string) eventDomain.Events
	GetEventsForTheWeek(date string) eventDomain.Events
	GetEventsForTheMonth(date string) eventDomain.Events
}

func NewServer(httpServerCfg *config.HTTPServer, app Application, logger *slog.Logger) *Server {
	return &Server{
		httpServerCfg,
		app,
		logger,
	}
}

func (s *Server) Start(ctx context.Context) error {
	router := NewRouter(s.app)
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
