package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MonkyMars/gecho/utils"
)

func TestCorrectMethodHandlers(t *testing.T) {
	// Test GET method
	t.Run(http.MethodGet, func(t *testing.T) {
		// Simulate a GET request
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/test", nil)

		handlers := NewHandlers()
		err := handlers.HandleMethod(w, r, http.MethodGet)
		if err != nil {
			t.Errorf("Expected nil error for allowed method, got %v", err)
		}

		if w.Result().StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
		}
	})

	// Test POST method
	t.Run(http.MethodPost, func(t *testing.T) {
		// Simulate a POST request
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/test", nil)

		handlers := NewHandlers()
		err := handlers.HandleMethod(w, r, http.MethodPost)
		if err != nil {
			t.Errorf("Expected nil error for allowed method, got %v", err)
		}

		if w.Result().StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
		}
	})

	// Test PUT method
	t.Run(http.MethodPut, func(t *testing.T) {
		// Simulate a PUT request
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPut, "/test", nil)

		handlers := NewHandlers()
		err := handlers.HandleMethod(w, r, http.MethodPut)
		if err != nil {
			t.Errorf("Expected nil error for allowed method, got %v", err)
		}

		if w.Result().StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
		}
	})

	// Test PATCH method
	t.Run(http.MethodPatch, func(t *testing.T) {
		// Simulate a PATCH request
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPatch, "/test", nil)

		handlers := NewHandlers()
		err := handlers.HandleMethod(w, r, http.MethodPatch)
		if err != nil {
			t.Errorf("Expected nil error for allowed method, got %v", err)
		}

		if w.Result().StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
		}
	})

	// Test DELETE method
	t.Run(http.MethodDelete, func(t *testing.T) {
		// Simulate a DELETE request
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodDelete, "/test", nil)

		handlers := NewHandlers()
		err := handlers.HandleMethod(w, r, http.MethodDelete)
		if err != nil {
			t.Errorf("Expected nil error for allowed method, got %v", err)
		}

		if w.Result().StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
		}
	})
}

func TestIncorrectMethodHandlers(t *testing.T) {
	// Test GET method with incorrect intended method
	t.Run(http.MethodGet, func(t *testing.T) {
		// Simulate a GET request
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/test", nil)

		handlers := NewHandlers()
		err := handlers.HandleMethod(w, r, http.MethodPost)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if w.Result().StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, w.Result().StatusCode)
		}

		var response utils.NewResponse
		if err := json.NewDecoder(w.Result().Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		expectedMsg := fmt.Sprintf("Method %s not allowed", http.MethodGet)
		if response.Message() != expectedMsg {
			t.Errorf("Expected message '%s', got '%s'", expectedMsg, response.Message())
		}
	})

	// Test POST method with incorrect intended method
	t.Run(http.MethodPost, func(t *testing.T) {
		// Simulate a POST request
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/test", nil)

		handlers := NewHandlers()
		err := handlers.HandleMethod(w, r, http.MethodGet)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if w.Result().StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, w.Result().StatusCode)
		}

		var response utils.NewResponse
		if err := json.NewDecoder(w.Result().Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		expectedMsg := fmt.Sprintf("Method %s not allowed", http.MethodPost)
		if response.Message() != expectedMsg {
			t.Errorf("Expected message '%s', got '%s'", expectedMsg, response.Message())
		}
	})

	t.Run(http.MethodPatch, func(t *testing.T) {
		// Simulate a PATCH request
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPatch, "/test", nil)

		handlers := NewHandlers()
		err := handlers.HandleMethod(w, r, http.MethodDelete)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if w.Result().StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, w.Result().StatusCode)
		}

		var response utils.NewResponse
		if err := json.NewDecoder(w.Result().Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		expectedMsg := fmt.Sprintf("Method %s not allowed", http.MethodPatch)
		if response.Message() != expectedMsg {
			t.Errorf("Expected message '%s', got '%s'", expectedMsg, response.Message())
		}
	})
}
