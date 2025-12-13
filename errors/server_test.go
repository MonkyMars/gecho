package errors

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MonkyMars/gecho/utils"
)

func TestInternalServerError(t *testing.T) {
	w := httptest.NewRecorder()
	err := InternalServerError(w, utils.Send())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, resp.StatusCode)
	}

	var response utils.NewResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Message() != utils.InternalServerErrorMessage {
		t.Errorf("Expected message '%s', got '%s'", utils.InternalServerErrorMessage, response.Message())
	}
}

func TestServiceUnavailable(t *testing.T) {
	w := httptest.NewRecorder()
	err := ServiceUnavailable(w, utils.Send())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("Expected status code %d, got %d", http.StatusServiceUnavailable, resp.StatusCode)
	}

	var response utils.NewResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Message() != utils.ServiceUnavailableMessage {
		t.Errorf("Expected message '%s', got '%s'", utils.ServiceUnavailableMessage, response.Message())
	}
}
