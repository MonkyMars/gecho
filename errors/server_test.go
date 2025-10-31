package errors

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MonkyMars/gecho/utils"
)

func TestInternalServerError(t *testing.T) {
	w := httptest.NewRecorder()
	rb := InternalServerError(w).WithMessage(utils.InternalServerErrorMessage)

	err := rb.Send()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, resp.StatusCode)
	}

	if rb.Response().Message() != utils.InternalServerErrorMessage {
		t.Errorf("Expected message '%s', got '%s'", utils.InternalServerErrorMessage, rb.Response().Message())
	}

	if rb.Response().Data() != nil {
		t.Errorf("Expected data to be nil for error responses, got '%v'", rb.Response().Data())
	}
}

func TestServiceUnavailable(t *testing.T) {
	w := httptest.NewRecorder()
	rb := ServiceUnavailable(w).WithMessage(utils.ServiceUnavailableMessage)

	err := rb.Send()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("Expected status code %d, got %d", http.StatusServiceUnavailable, resp.StatusCode)
	}

	if rb.Response().Message() != utils.ServiceUnavailableMessage {
		t.Errorf("Expected message '%s', got '%s'", utils.ServiceUnavailableMessage, rb.Response().Message())
	}

	if rb.Response().Data() != nil {
		t.Errorf("Expected data to be nil for error responses, got '%v'", rb.Response().Data())
	}
}
