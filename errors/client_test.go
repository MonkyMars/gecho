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
	BadRequest(w, utils.WithData(map[string]string{"field": "invalid"}), utils.Send())

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
	Unauthorized(w, utils.Send())

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
	Forbidden(w, utils.Send())

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
	NotFound(w, utils.Send())

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
	MethodNotAllowed(w, utils.Send())

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

func TestGone(t *testing.T) {
	w := httptest.NewRecorder()
	Gone(w, utils.Send())

	resp := w.Result()
	if resp.StatusCode != http.StatusGone {
		t.Errorf("Expected status code %d, got %d", http.StatusGone, resp.StatusCode)
	}

	var response utils.NewResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Message() != utils.GoneMessage {
		t.Errorf("Expected message '%s', got '%s'", utils.GoneMessage, response.Message())
	}

	if response.Success() {
		t.Errorf("Expected success to be false for Gone response")
	}
}

func TestUnprocessableEntity(t *testing.T) {
	w := httptest.NewRecorder()
	validationErrors := map[string]string{"email": "invalid format"}
	UnprocessableEntity(w, utils.WithData(validationErrors), utils.Send())

	resp := w.Result()
	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("Expected status code %d, got %d", http.StatusUnprocessableEntity, resp.StatusCode)
	}

	var response utils.NewResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Message() != utils.UnprocessableEntityMessage {
		t.Errorf("Expected message '%s', got '%s'", utils.UnprocessableEntityMessage, response.Message())
	}

	dataMap, ok := response.Data().(map[string]any)
	if !ok || dataMap["email"] != "invalid format" {
		t.Errorf("Expected data map with email 'invalid format', got '%v'", response.Data())
	}
}
