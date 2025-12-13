package errors

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MonkyMars/gecho/utils"
)

func TestBadRequest(t *testing.T) {
	w := httptest.NewRecorder()
	err := BadRequest(w, utils.WithData(map[string]string{"field": "invalid"}), utils.Send())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	var response utils.NewResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Message() != utils.BadRequestMessage {
		t.Errorf("Expected message '%s', got '%s'", utils.BadRequestMessage, response.Message())
	}

	if !response.Success() {
		// Expected for error responses
	}
}

func TestUnauthorized(t *testing.T) {
	w := httptest.NewRecorder()
	err := Unauthorized(w, utils.Send())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, resp.StatusCode)
	}

	var response utils.NewResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Message() != utils.UnauthorizedMessage {
		t.Errorf("Expected message '%s', got '%s'", utils.UnauthorizedMessage, response.Message())
	}
}

func TestForbidden(t *testing.T) {
	w := httptest.NewRecorder()
	err := Forbidden(w, utils.Send())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusForbidden {
		t.Errorf("Expected status code %d, got %d", http.StatusForbidden, resp.StatusCode)
	}

	var response utils.NewResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Message() != utils.ForbiddenMessage {
		t.Errorf("Expected message '%s', got '%s'", utils.ForbiddenMessage, response.Message())
	}
}

func TestNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	err := NotFound(w, utils.Send())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, resp.StatusCode)
	}

	var response utils.NewResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Message() != utils.NotFoundMessage {
		t.Errorf("Expected message '%s', got '%s'", utils.NotFoundMessage, response.Message())
	}
}

func TestMethodNotAllowed(t *testing.T) {
	w := httptest.NewRecorder()
	err := MethodNotAllowed(w, utils.Send())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}

	var response utils.NewResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Message() != utils.MethodNotAllowedMessage {
		t.Errorf("Expected message '%s', got '%s'", utils.MethodNotAllowedMessage, response.Message())
	}
}
