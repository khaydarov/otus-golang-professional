package middleware

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

var RequestIDHeader = "X-Request-Id"
var RequestIDKey = "requestID"

func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get(RequestIDHeader)
		if requestID == "" {
			id, err := uuid.NewV7()
			if err != nil {
				panic("could not generate request id")
			}

			requestID = id.String()
		}
		ctx = context.WithValue(ctx, RequestIDKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
		return reqID
	}
	return ""
}
