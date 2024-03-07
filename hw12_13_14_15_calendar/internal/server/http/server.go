package internalhttp

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	"log/slog"
	"net/http"
)

type Server struct {
	cfg *config.HTTPServer
	app Application
	log *slog.Logger
}

type Application interface {
	CreateEvent(ctx context.Context, title string) (string, error)
}

func NewServer(httpServerCfg *config.HTTPServer, logger *slog.Logger, app Application) *Server {
	return &Server{
		httpServerCfg,
		app,
		logger,
	}
}

func (s *Server) Start(ctx context.Context) error {
	router := chi.NewRouter()

	router.Use(LoggerMiddleware("./logs/log.txt"))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	server := &http.Server{
		Addr:        fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port),
		ReadTimeout: 0,
		Handler:     router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			s.log.Error("failed to start server: %s", err)
		}
	}()

	<-ctx.Done()
	return nil
}

func (s *Server) Stop(_ context.Context) error {
	return nil
}
