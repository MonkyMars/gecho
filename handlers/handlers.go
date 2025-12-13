package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MonkyMars/gecho/errors"
	"github.com/MonkyMars/gecho/utils"
)

type Handlers struct{}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) HandleMethod(w http.ResponseWriter, r *http.Request, intendedMethod string) error {
	if r.Method != intendedMethod {
		return errors.MethodNotAllowed(w, utils.Send(), utils.WithMessage(fmt.Sprintf("Method %s not allowed", r.Method)))
	}
	return nil
}

func (h *Handlers) CreateLoggingMiddleware(logger *utils.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return h.HandleLogging(next, logger)
	}
}

func (h *Handlers) HandleLogging(next http.Handler, logger *utils.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Create a response writer wrapper to capture status code
		wrapper := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(wrapper, r)

		duration := time.Since(start)
		if wrapper.statusCode >= 500 {
			logger.Error(
				utils.Field("method", r.Method),
				utils.Field("path", r.URL.Path),
				utils.Field("status", wrapper.statusCode),
				utils.Field("duration", duration),
				utils.Field("remote_addr", r.RemoteAddr),
			)

		} else if wrapper.statusCode >= 400 {
			logger.Warn(
				utils.Field("method", r.Method),
				utils.Field("path", r.URL.Path),
				utils.Field("status", wrapper.statusCode),
				utils.Field("duration", duration),
				utils.Field("remote_addr", r.RemoteAddr),
			)
		} else {
			logger.Info(
				utils.Field("method", r.Method),
				utils.Field("path", r.URL.Path),
				utils.Field("status", wrapper.statusCode),
				utils.Field("duration", duration),
				utils.Field("remote_addr", r.RemoteAddr),
			)
		}
	})
}

// responseWriter is a wrapper to capture the status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
