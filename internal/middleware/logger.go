package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

func NewMiddlewareLogger(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log := logger.With(
			slog.String("component", "middleware_logger"),
		)

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			entry := log.With(
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				// Если нужно логировать id запросов:
				// slog.String("request_id", middleware.GetReqID(r.Context())),
			)

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			start := time.Now()

			defer func() {
				status := ww.Status()
				if status == 0 {
					status = http.StatusOK
				}

				args := []any{
					slog.Int("status", status),
					slog.Int("bytes", ww.BytesWritten()),
					slog.Duration("duration", time.Since(start)),
				}

				switch {
				case status >= 500:
					entry.Error("request completed", args...)
				case status >= 400:
					entry.Warn("request completed", args...)
				default:
					entry.Info("request completed", args...)
				}
			}()

			next.ServeHTTP(ww, r)
		})
	}
}
