package logger

import (
	"log/slog"
	"os"
)

const (
	Debug   = "debug"
	Info    = "info"
	Warning = "warning"
	Error   = "error"
)

func New(logLevel string) *slog.Logger {
	var level slog.Level
	switch logLevel {
	case Debug:
		level = slog.LevelDebug
	case Info:
		level = slog.LevelInfo
	case Warning:
		level = slog.LevelWarn
	case Error:
		level = slog.LevelError
	default:
		level = slog.LevelDebug
	}

	return slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}),
	)
}
