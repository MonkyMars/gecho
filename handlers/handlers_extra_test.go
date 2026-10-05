package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MonkyMars/gecho/utils"
)

func TestCreateLoggingMiddleware(t *testing.T) {
	var output bytes.Buffer
	logger := utils.NewLogger(utils.NewConfig(
		utils.WithShowCaller(false),
		utils.WithOutput(&output),
		utils.WithErrorOutput(&output),
	))
	handler := NewHandlers().CreateLoggingMiddleware(logger)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/health", nil))
	if w.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusNoContent)
	}
}
