package httplogger

import (
	"context"
	"io"
	"log"
	"log/slog"
)

type HTTPRequestLoggerHandlerOptions struct {
	Level slog.Leveler
}

type HTTPRequestLoggerHandler struct {
	slog.Handler

	opts  HTTPRequestLoggerHandlerOptions
	l     *log.Logger
	attrs []slog.Attr
}

func NewHTTPRequestLoggerHandler(out io.Writer, opts HTTPRequestLoggerHandlerOptions) *HTTPRequestLoggerHandler {
	return &HTTPRequestLoggerHandler{
		opts: opts,
		l:    log.New(out, "", 0),
	}
}

func (h *HTTPRequestLoggerHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

func (h *HTTPRequestLoggerHandler) Handle(_ context.Context, r slog.Record) error {
	h.l.Printf(r.Message)
	return nil
}

func (h *HTTPRequestLoggerHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &HTTPRequestLoggerHandler{
		Handler: h.Handler,
		l:       h.l,
		attrs:   attrs,
	}
}

func (h *HTTPRequestLoggerHandler) WithGroup(name string) slog.Handler {
	return &HTTPRequestLoggerHandler{
		Handler: h.Handler.WithGroup(name),
		l:       h.l,
	}
}
