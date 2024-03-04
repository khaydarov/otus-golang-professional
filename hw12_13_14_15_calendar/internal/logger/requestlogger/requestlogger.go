package requestlogger

import (
	"context"
	"io"
	stdLog "log"
	"log/slog"
)

type HttpRequestLoggerHandlerOptions struct {
	Level slog.Leveler
}

type HttpRequestLoggerHandler struct {
	slog.Handler

	opts  HttpRequestLoggerHandlerOptions
	l     *stdLog.Logger
	attrs []slog.Attr
}

func NewHttpRequestLoggerHandler(out io.Writer, opts HttpRequestLoggerHandlerOptions) *HttpRequestLoggerHandler {
	return &HttpRequestLoggerHandler{
		opts: opts,
		l:    stdLog.New(out, "", 0),
	}
}

func (h *HttpRequestLoggerHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

func (h *HttpRequestLoggerHandler) Handle(_ context.Context, r slog.Record) error {
	h.l.Printf(r.Message)
	return nil
}

func (h *HttpRequestLoggerHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &HttpRequestLoggerHandler{
		Handler: h.Handler,
		l:       h.l,
		attrs:   attrs,
	}
}

func (h *HttpRequestLoggerHandler) WithGroup(name string) slog.Handler {
	return &HttpRequestLoggerHandler{
		Handler: h.Handler.WithGroup(name),
		l:       h.l,
	}
}
