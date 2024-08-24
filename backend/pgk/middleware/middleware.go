package middleware

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

func CreateCorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json;")
		w.Header().Set("access-control-allow-origin", "*")
		w.Header().Set("access-control-allow-methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("access-control-allow-headers", "*")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CreateLoggerMiddleware(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId, _ := uuid.NewUUID()

			log := logger.With("requestId", requestId.String())

			log.Debug("Call to endpoint",
				"method", r.Method,
				"path", r.URL.EscapedPath(),
			)

			ctx := context.WithValue(r.Context(), "logger", log)

			next.ServeHTTP(w, r.WithContext(ctx))

			log.Debug("Finished call to endpoint",
				"method", r.Method,
				"path", r.URL.EscapedPath(),
			)
		})
	}
}
