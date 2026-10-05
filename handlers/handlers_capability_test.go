package handlers

import (
	"bufio"
	"errors"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MonkyMars/gecho/utils"
)

func TestResponseWriterPreservesOptionalCapabilities(t *testing.T) {
	var flushed bool
	base := &capabilityWriter{ResponseRecorder: httptest.NewRecorder(), flushed: &flushed}
	logger := utils.NewLogger(utils.NewConfig(
		utils.WithShowCaller(false),
		utils.WithOutput(ioDiscard{}),
		utils.WithErrorOutput(ioDiscard{}),
	))

	handler := NewHandlers().HandleLogging(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := http.NewResponseController(w).Flush(); err != nil {
			t.Fatalf("Flush() error: %v", err)
		}
		if _, _, err := w.(http.Hijacker).Hijack(); !errors.Is(err, errHijack) {
			t.Fatalf("Hijack() error = %v, want %v", err, errHijack)
		}
		w.WriteHeader(http.StatusNoContent)
	}), logger)

	handler.ServeHTTP(base, httptest.NewRequest(http.MethodGet, "/", nil))
	if !flushed {
		t.Fatal("Flush was not forwarded to the underlying writer")
	}
}

var errHijack = errors.New("hijack unavailable")

type capabilityWriter struct {
	*httptest.ResponseRecorder
	flushed *bool
}

func (w *capabilityWriter) Flush() {
	*w.flushed = true
}

func (w *capabilityWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return nil, nil, errHijack
}

type ioDiscard struct{}

func (ioDiscard) Write(p []byte) (int, error) { return len(p), nil }
