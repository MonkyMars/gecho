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

func (h *Handlers) HandleMethod(w http.ResponseWriter, r *http.Request, intendedMethod string) *utils.ResponseBuilder {
	method := r.Method
	if method != intendedMethod {
		return errors.MethodNotAllowed(w).WithMessage(fmt.Sprintf("Method %s not allowed", method))
	}

	return nil
}

func (h *Handlers) HandleLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Create a response writer wrapper to capture status code
		wrapper := &responseWriter{ResponseWriter: w, statusCode: r.Response.StatusCode}
		next.ServeHTTP(wrapper, r)

		logger := utils.NewDefaultLogger()

		duration := time.Since(start)
		if wrapper.statusCode >= 500 {
			logger.Error(
				"method", r.Method,
				"path", r.URL.Path,
				"status", wrapper.statusCode,
				"duration", duration,
				"remote_addr", r.RemoteAddr,
			)
		} else if wrapper.statusCode >= 400 {
			logger.Warn(
				"method", r.Method,
				"path", r.URL.Path,
				"status", wrapper.statusCode,
				"duration", duration,
				"remote_addr", r.RemoteAddr,
			)
		} else {
			logger.Info(
				"method", r.Method,
				"path", r.URL.Path,
				"status", wrapper.statusCode,
				"duration", duration,
				"remote_addr", r.RemoteAddr,
			)
		}
	})
}

// responseWriter is a wrapper to capture the status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}
