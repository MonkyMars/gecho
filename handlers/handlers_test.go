package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCorrectMethodHandlers(t *testing.T) {
	// Test GET method
	t.Run(http.MethodGet, func(t *testing.T) {
		// Simulate a GET request
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/test", nil)

		handlers := NewHandlers()
		responseBuilder := handlers.HandleMethod(w, r, http.MethodGet)
		if responseBuilder != nil {
			t.Errorf("Expected nil response for allowed method, got %v", responseBuilder)
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
		responseBuilder := handlers.HandleMethod(w, r, http.MethodGet)
		if responseBuilder == nil {
			t.Errorf("Expected non-nil response for disallowed method, got nil")
		} else {
			err := responseBuilder.Send()
			if err != nil {
				t.Errorf("Expected no error on Send(), got %v", err)
			}

			if w.Result().StatusCode != http.StatusMethodNotAllowed {
				t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, w.Result().StatusCode)
			}
		}
	})

	// Test PUT method
	t.Run(http.MethodPut, func(t *testing.T) {
		// Simulate a PUT request
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPut, "/test", nil)

		handlers := NewHandlers()
		responseBuilder := handlers.HandleMethod(w, r, http.MethodPut)
		if responseBuilder != nil {
			t.Errorf("Expected nil response for allowed method, got %v", responseBuilder)
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
		responseBuilder := handlers.HandleMethod(w, r, http.MethodPatch)
		if responseBuilder != nil {
			t.Errorf("Expected nil response for allowed method, got %v", responseBuilder)
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
		responseBuilder := handlers.HandleMethod(w, r, http.MethodDelete)
		if responseBuilder != nil {
			t.Errorf("Expected nil response for allowed method, got %v", responseBuilder)
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
		responseBuilder := handlers.HandleMethod(w, r, http.MethodPost)
		if responseBuilder == nil {
			t.Errorf("Expected non-nil response for disallowed method, got nil")
		} else {
			err := responseBuilder.Send()
			if err != nil {
				t.Errorf("Expected no error on Send(), got %v", err)
			}

			if w.Result().StatusCode != http.StatusMethodNotAllowed {
				t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, w.Result().StatusCode)
			}

			if !strings.Contains(w.Body.String(), fmt.Sprintf("Method %s not allowed", http.MethodGet)) {
				t.Errorf("Expected message 'Method GET not allowed' to be present, got '%s'", w.Body.String())
			}
		}
	})

	// Test POST method with incorrect intended method
	t.Run(http.MethodPost, func(t *testing.T) {
		// Simulate a POST request
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/test", nil)

		handlers := NewHandlers()
		responseBuilder := handlers.HandleMethod(w, r, http.MethodGet)
		if responseBuilder == nil {
			t.Errorf("Expected non-nil response for disallowed method, got nil")
		} else {
			err := responseBuilder.Send()
			if err != nil {
				t.Errorf("Expected no error on Send(), got %v", err)
			}

			if w.Result().StatusCode != http.StatusMethodNotAllowed {
				t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, w.Result().StatusCode)
			}

			if !strings.Contains(w.Body.String(), fmt.Sprintf("Method %s not allowed", http.MethodPost)) {
				t.Errorf("Expected message 'Method POST not allowed' to be present, got '%s'", w.Body.String())
			}
		}
	})

	t.Run(http.MethodPatch, func(t *testing.T) {
		// Simulate a PATCH request
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPatch, "/test", nil)

		handlers := NewHandlers()
		responseBuilder := handlers.HandleMethod(w, r, http.MethodDelete)
		if responseBuilder == nil {
			t.Errorf("Expected non-nil response for disallowed method, got nil")
		} else {
			err := responseBuilder.Send()
			if err != nil {
				t.Errorf("Expected no error on Send(), got %v", err)
			}

			if w.Result().StatusCode != http.StatusMethodNotAllowed {
				t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, w.Result().StatusCode)
			}

			if !strings.Contains(w.Body.String(), fmt.Sprintf("Method %s not allowed", http.MethodPatch)) {
				t.Errorf("Expected message 'Method PATCH not allowed' to be present, got '%s'", w.Body.String())
			}
		}
	})
}
