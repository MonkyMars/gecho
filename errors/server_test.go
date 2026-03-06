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
	InternalServerError(w, utils.Send())

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
	ServiceUnavailable(w, utils.Send())

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

func TestNotImplemented(t *testing.T) {
	w := httptest.NewRecorder()
	NotImplemented(w, utils.Send())

	resp := w.Result()
	if resp.StatusCode != http.StatusNotImplemented {
		t.Errorf("Expected status code %d, got %d", http.StatusNotImplemented, resp.StatusCode)
	}

	var response utils.NewResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Message() != utils.NotImplementedMessage {
		t.Errorf("Expected message '%s', got '%s'", utils.NotImplementedMessage, response.Message())
	}

	if response.Success() {
		t.Errorf("Expected success to be false for NotImplemented response")
	}
}

func TestBadGateway(t *testing.T) {
	w := httptest.NewRecorder()
	BadGateway(w, utils.Send())

	resp := w.Result()
	if resp.StatusCode != http.StatusBadGateway {
		t.Errorf("Expected status code %d, got %d", http.StatusBadGateway, resp.StatusCode)
	}

	var response utils.NewResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Message() != utils.BadGatewayMessage {
		t.Errorf("Expected message '%s', got '%s'", utils.BadGatewayMessage, response.Message())
	}
}

func TestGatewayTimeout(t *testing.T) {
	w := httptest.NewRecorder()
	GatewayTimeout(w, utils.Send())

	resp := w.Result()
	if resp.StatusCode != http.StatusGatewayTimeout {
		t.Errorf("Expected status code %d, got %d", http.StatusGatewayTimeout, resp.StatusCode)
	}

	var response utils.NewResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Message() != utils.GatewayTimeoutMessage {
		t.Errorf("Expected message '%s', got '%s'", utils.GatewayTimeoutMessage, response.Message())
	}
}
