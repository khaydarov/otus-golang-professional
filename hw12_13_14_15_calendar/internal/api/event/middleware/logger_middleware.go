package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/khaydarov/otus-golang-professional/hw12_13_14_15_calendar/internal/logger/httplogger"
)

func LoggerMiddleware(filePath string) func(next http.Handler) http.Handler {
	logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}

	log := slog.New(
		httplogger.NewHTTPRequestLoggerHandler(
			logFile,
			httplogger.HTTPRequestLoggerHandlerOptions{Level: slog.LevelInfo},
		),
	)

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()
			defer func() {
				message := fmt.Sprintf(
					"%s [%s] %s %s %s %d %s %s\n",
					r.RemoteAddr,
					time.Now().Format("02/Jan/2006:15:04:05 -0700"),
					r.Method,
					r.URL.Path,
					r.Proto,
					ww.Status(),
					time.Since(t1).String(),
					r.UserAgent(),
				)

				log.Log(context.Background(), slog.LevelInfo, message)
			}()

			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}
