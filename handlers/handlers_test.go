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

func TestHandleLogging(t *testing.T) {
	t.Run("Success_200", func(t *testing.T) {
		// Create a test handler that returns 200
		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})

		handlers := NewHandlers()
		loggingHandler := handlers.HandleLogging(testHandler, utils.NewDefaultLogger())

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/test", nil)
		r.RemoteAddr = "127.0.0.1:12345"

		loggingHandler.ServeHTTP(w, r)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}

		if w.Body.String() != "OK" {
			t.Errorf("Expected body 'OK', got '%s'", w.Body.String())
		}
	})

	t.Run("Success_201", func(t *testing.T) {
		// Create a test handler that returns 201
		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Created"))
		})

		handlers := NewHandlers()
		loggingHandler := handlers.HandleLogging(testHandler, utils.NewDefaultLogger())

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/api/users", nil)

		loggingHandler.ServeHTTP(w, r)

		if w.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
		}
	})

	t.Run("ClientError_400", func(t *testing.T) {
		// Create a test handler that returns 400
		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
		})

		handlers := NewHandlers()
		loggingHandler := handlers.HandleLogging(testHandler, utils.NewDefaultLogger())

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/api/users", nil)

		loggingHandler.ServeHTTP(w, r)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("ClientError_404", func(t *testing.T) {
		// Create a test handler that returns 404
		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not Found"))
		})

		handlers := NewHandlers()
		loggingHandler := handlers.HandleLogging(testHandler, utils.NewDefaultLogger())

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/api/users/999", nil)

		loggingHandler.ServeHTTP(w, r)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
		}
	})

	t.Run("ClientError_401", func(t *testing.T) {
		// Create a test handler that returns 401
		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		})

		handlers := NewHandlers()
		loggingHandler := handlers.HandleLogging(testHandler, utils.NewDefaultLogger())

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/api/protected", nil)

		loggingHandler.ServeHTTP(w, r)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, w.Code)
		}
	})

	t.Run("ServerError_500", func(t *testing.T) {
		// Create a test handler that returns 500
		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
		})

		handlers := NewHandlers()
		loggingHandler := handlers.HandleLogging(testHandler, utils.NewDefaultLogger())

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/api/error", nil)

		loggingHandler.ServeHTTP(w, r)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
		}
	})

	t.Run("ServerError_503", func(t *testing.T) {
		// Create a test handler that returns 503
		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("Service Unavailable"))
		})

		handlers := NewHandlers()
		loggingHandler := handlers.HandleLogging(testHandler, utils.NewDefaultLogger())

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/api/service", nil)

		loggingHandler.ServeHTTP(w, r)

		if w.Code != http.StatusServiceUnavailable {
			t.Errorf("Expected status code %d, got %d", http.StatusServiceUnavailable, w.Code)
		}
	})

	t.Run("DefaultStatusCode", func(t *testing.T) {
		// Create a test handler that doesn't explicitly set status code
		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Response without explicit status"))
		})

		handlers := NewHandlers()
		loggingHandler := handlers.HandleLogging(testHandler, utils.NewDefaultLogger())

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/test", nil)

		loggingHandler.ServeHTTP(w, r)

		// Default status should be 200 OK
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
	})

	t.Run("DifferentHTTPMethods", func(t *testing.T) {
		methods := []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		}

		for _, method := range methods {
			t.Run(method, func(t *testing.T) {
				testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				})

				handlers := NewHandlers()
				loggingHandler := handlers.HandleLogging(testHandler, utils.NewDefaultLogger())

				w := httptest.NewRecorder()
				r := httptest.NewRequest(method, "/test", nil)

				loggingHandler.ServeHTTP(w, r)

				if w.Code != http.StatusOK {
					t.Errorf("Expected status code %d for method %s, got %d", http.StatusOK, method, w.Code)
				}
			})
		}
	})

	t.Run("PreservesResponseBody", func(t *testing.T) {
		expectedBody := `{"message": "test response", "data": [1, 2, 3]}`
		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(expectedBody))
		})

		handlers := NewHandlers()
		loggingHandler := handlers.HandleLogging(testHandler, utils.NewDefaultLogger())

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/api/data", nil)

		loggingHandler.ServeHTTP(w, r)

		if w.Body.String() != expectedBody {
			t.Errorf("Expected body '%s', got '%s'", expectedBody, w.Body.String())
		}

		if w.Header().Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type 'application/json', got '%s'", w.Header().Get("Content-Type"))
		}
	})

	t.Run("PreservesHeaders", func(t *testing.T) {
		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Custom-Header", "CustomValue")
			w.Header().Set("Cache-Control", "no-cache")
			w.WriteHeader(http.StatusOK)
		})

		handlers := NewHandlers()
		loggingHandler := handlers.HandleLogging(testHandler, utils.NewDefaultLogger())

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/test", nil)

		loggingHandler.ServeHTTP(w, r)

		if w.Header().Get("X-Custom-Header") != "CustomValue" {
			t.Errorf("Expected X-Custom-Header 'CustomValue', got '%s'", w.Header().Get("X-Custom-Header"))
		}

		if w.Header().Get("Cache-Control") != "no-cache" {
			t.Errorf("Expected Cache-Control 'no-cache', got '%s'", w.Header().Get("Cache-Control"))
		}
	})

	t.Run("CapturesRequestPath", func(t *testing.T) {
		testPaths := []string{
			"/",
			"/api/users",
			"/api/users/123",
			"/api/users/123/posts",
			"/health",
		}

		for _, path := range testPaths {
			t.Run(path, func(t *testing.T) {
				testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if r.URL.Path != path {
						t.Errorf("Expected path '%s', got '%s'", path, r.URL.Path)
					}
					w.WriteHeader(http.StatusOK)
				})

				handlers := NewHandlers()
				loggingHandler := handlers.HandleLogging(testHandler, utils.NewDefaultLogger())

				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodGet, path, nil)

				loggingHandler.ServeHTTP(w, r)
			})
		}
	})
}
