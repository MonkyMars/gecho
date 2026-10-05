package handlers

import (
	"bufio"
	"fmt"
	"net"
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
		if err := errors.MethodNotAllowed(w, utils.WithMessage(fmt.Sprintf("Method %s not allowed", r.Method))).Send(); err != nil {
			return err
		}
		return fmt.Errorf("method %s not allowed", r.Method)
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
	statusCode  int
	wroteHeader bool
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}
	rw.statusCode = code
	rw.wroteHeader = true
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(body []byte) (int, error) {
	if rw.statusCode == http.StatusOK {
		rw.WriteHeader(http.StatusOK)
	}
	return rw.ResponseWriter.Write(body)
}

func (rw *responseWriter) Unwrap() http.ResponseWriter {
	return rw.ResponseWriter
}

func (rw *responseWriter) Flush() {
	flusher, ok := rw.ResponseWriter.(http.Flusher)
	if !ok {
		return
	}
	if rw.statusCode == http.StatusOK {
		rw.WriteHeader(http.StatusOK)
	}
	flusher.Flush()
}

func (rw *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := rw.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("http.ResponseWriter does not support hijacking")
	}
	return hijacker.Hijack()
}

func (rw *responseWriter) Push(target string, opts *http.PushOptions) error {
	pusher, ok := rw.ResponseWriter.(http.Pusher)
	if !ok {
		return http.ErrNotSupported
	}
	return pusher.Push(target, opts)
}
