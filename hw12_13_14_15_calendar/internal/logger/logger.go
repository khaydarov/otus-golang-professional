package logger

import (
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envStage = "stage"
	envProd  = "prod"
)

func New(env string) *slog.Logger {
	var level slog.Level
	switch env {
	case envLocal:
		level = slog.LevelDebug
	case envStage:
		level = slog.LevelDebug
	case envProd:
		level = slog.LevelInfo
	default:
		level = slog.LevelDebug
	}

	return slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}),
	)
}
