package internalhttp

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
)

type Server struct {
	Config *config.HTTPServer
	*slog.Logger
	Application
}

type Application interface{}

func NewServer(httpServerConfig *config.HTTPServer, logger *slog.Logger, app Application) *Server {
	return &Server{
		httpServerConfig,
		logger,
		app,
	}
}

func (s *Server) Start(ctx context.Context) error {
	router := chi.NewRouter()
	router.Use(LoggerMiddleware("../../log.txt"))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port),
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			s.Logger.Error("failed to start server: %s", err)
		}
	}()

	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return nil
}
