package middlewares

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

// middleware based on https://github.com/igknot/chi-zap-ecs-logger

type chilogger struct {
	logZ *zap.SugaredLogger
	name string
}

func NewZapMiddleware(name string, logger *zap.SugaredLogger) func(next http.Handler) http.Handler {
	return chilogger{
		logZ: logger,
		name: name,
	}.middleware
}

func (c chilogger) middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var requestID string
		if reqID := r.Context().Value(middleware.RequestIDKey); reqID != nil {
			requestID = reqID.(string)
		} else {
			requestID = r.Header.Get("X-Request-Id")
		}
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)

		latency := time.Since(start)

		if c.logZ != nil {
			c.logZ.Infow(
				"request",
				"request_id", requestID,
				"method", r.Method,
				"uri", r.RequestURI,
				"status", ww.Status(),
				"took", latency,
				"remote_addr", r.RemoteAddr,
				"user_agent", r.UserAgent(),
			)
		}

	}
	return http.HandlerFunc(fn)
}
