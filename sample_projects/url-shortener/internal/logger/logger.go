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

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envStage:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	//opts := slogpretty.PrettyHandlerOptions{
	//	SlogOpts: &slog.HandlerOptions{
	//		Level: slog.LevelDebug,
	//	},
	//}
	//
	//handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
	) // slog.New(handler)
}
