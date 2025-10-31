package errors

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MonkyMars/gecho/utils"
)

func TestBadRequest(t *testing.T) {
	w := httptest.NewRecorder()
	rb := BadRequest(w).WithMessage(utils.BadRequestMessage).WithData(map[string]string{"field": "invalid"})

	err := rb.Send()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	if rb.Response().Message() != utils.BadRequestMessage {
		t.Errorf("Expected message '%s', got '%s'", utils.BadRequestMessage, rb.Response().Message())
	}

	dataMap, ok := rb.Response().Data().(map[string]string)
	if !ok || dataMap["field"] != "invalid" {
		t.Errorf("Expected data map with field 'invalid', got '%v'", rb.Response().Data())
	}
}

func TestUnauthorized(t *testing.T) {
	w := httptest.NewRecorder()
	rb := Unauthorized(w).WithMessage(utils.UnauthorizedMessage)

	err := rb.Send()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, resp.StatusCode)
	}

	if rb.Response().Message() != utils.UnauthorizedMessage {
		t.Errorf("Expected message '%s', got '%s'", utils.UnauthorizedMessage, rb.Response().Message())
	}
}

func TestForbidden(t *testing.T) {
	w := httptest.NewRecorder()
	rb := Forbidden(w).WithMessage(utils.ForbiddenMessage)

	err := rb.Send()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusForbidden {
		t.Errorf("Expected status code %d, got %d", http.StatusForbidden, resp.StatusCode)
	}

	if rb.Response().Message() != utils.ForbiddenMessage {
		t.Errorf("Expected message '%s', got '%s'", utils.ForbiddenMessage, rb.Response().Message())
	}
}

func TestNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	rb := NotFound(w).WithMessage(utils.NotFoundMessage)

	err := rb.Send()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, resp.StatusCode)
	}

	if rb.Response().Message() != utils.NotFoundMessage {
		t.Errorf("Expected message '%s', got '%s'", utils.NotFoundMessage, rb.Response().Message())
	}
}

func TestMethodNotAllowed(t *testing.T) {
	w := httptest.NewRecorder()
	rb := MethodNotAllowed(w).WithMessage(utils.MethodNotAllowedMessage)

	err := rb.Send()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}

	if rb.Response().Message() != utils.MethodNotAllowedMessage {
		t.Errorf("Expected message '%s', got '%s'", utils.MethodNotAllowedMessage, rb.Response().Message())
	}
}
