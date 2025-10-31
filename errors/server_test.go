package errors

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MonkyMars/gecho/utils"
)

func TestInternalServerError(t *testing.T) {
	w := httptest.NewRecorder()
	rb := InternalServerError(w).WithMessage(utils.INTERNAL_SERVER_ERROR_MESSAGE)

	err := rb.Send()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, resp.StatusCode)
	}

	if rb.Response.Message != utils.INTERNAL_SERVER_ERROR_MESSAGE {
		t.Errorf("Expected message '%s', got '%s'", utils.INTERNAL_SERVER_ERROR_MESSAGE, rb.Response.Message)
	}

	if rb.Response.Data != nil {
		t.Errorf("Expected data to be nil for error responses, got '%v'", rb.Response.Data)
	}
}

func TestServiceUnavailable(t *testing.T) {
	w := httptest.NewRecorder()
	rb := ServiceUnavailable(w).WithMessage(utils.SERVICE_UNAVAILABLE_MESSAGE)

	err := rb.Send()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("Expected status code %d, got %d", http.StatusServiceUnavailable, resp.StatusCode)
	}

	if rb.Response.Message != utils.SERVICE_UNAVAILABLE_MESSAGE {
		t.Errorf("Expected message '%s', got '%s'", utils.SERVICE_UNAVAILABLE_MESSAGE, rb.Response.Message)
	}

	if rb.Response.Data != nil {
		t.Errorf("Expected data to be nil for error responses, got '%v'", rb.Response.Data)
	}
}
