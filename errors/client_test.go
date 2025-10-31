package errors

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MonkyMars/gecho/utils"
)

func TestBadRequest(t *testing.T) {
	w := httptest.NewRecorder()
	rb := BadRequest(w).WithMessage(utils.BAD_REQUEST_MESSAGE).WithData(map[string]string{"field": "invalid"})

	err := rb.Send()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	if rb.Response.Message != utils.BAD_REQUEST_MESSAGE {
		t.Errorf("Expected message '%s', got '%s'", utils.BAD_REQUEST_MESSAGE, rb.Response.Message)
	}

	dataMap, ok := rb.Response.Data.(map[string]string)
	if !ok || dataMap["field"] != "invalid" {
		t.Errorf("Expected data map with field 'invalid', got '%v'", rb.Response.Data)
	}
}

func TestUnauthorized(t *testing.T) {
	w := httptest.NewRecorder()
	rb := Unauthorized(w).WithMessage(utils.UNAUTHORIZED_MESSAGE)

	err := rb.Send()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, resp.StatusCode)
	}

	if rb.Response.Message != utils.UNAUTHORIZED_MESSAGE {
		t.Errorf("Expected message '%s', got '%s'", utils.UNAUTHORIZED_MESSAGE, rb.Response.Message)
	}
}

func TestForbidden(t *testing.T) {
	w := httptest.NewRecorder()
	rb := Forbidden(w).WithMessage(utils.FORBIDDEN_MESSAGE)

	err := rb.Send()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusForbidden {
		t.Errorf("Expected status code %d, got %d", http.StatusForbidden, resp.StatusCode)
	}

	if rb.Response.Message != utils.FORBIDDEN_MESSAGE {
		t.Errorf("Expected message '%s', got '%s'", utils.FORBIDDEN_MESSAGE, rb.Response.Message)
	}
}

func TestNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	rb := NotFound(w).WithMessage(utils.NOT_FOUND_MESSAGE)

	err := rb.Send()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, resp.StatusCode)
	}

	if rb.Response.Message != utils.NOT_FOUND_MESSAGE {
		t.Errorf("Expected message '%s', got '%s'", utils.NOT_FOUND_MESSAGE, rb.Response.Message)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	w := httptest.NewRecorder()
	rb := MethodNotAllowed(w).WithMessage(utils.METHOD_NOT_ALLOWED_MESSAGE)

	err := rb.Send()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}

	if rb.Response.Message != utils.METHOD_NOT_ALLOWED_MESSAGE {
		t.Errorf("Expected message '%s', got '%s'", utils.METHOD_NOT_ALLOWED_MESSAGE, rb.Response.Message)
	}
}