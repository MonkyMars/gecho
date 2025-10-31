package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestExtractResponseBody(t *testing.T) {
	w := httptest.NewRecorder()
	rb := NewOK(w).
		WithMessage("Extract Test").
		WithData(map[string]string{"extract": "test"}).
		WithStatus(http.StatusOK)

	err := rb.Send()
	if err != nil {
		t.Errorf("Expected no error on Send(), got %v", err)
	}

	resp := w.Result()

	// Extract response body and verify
	val, err := ExtractResponseBody[NewResponse](resp)
	if err != nil {
		t.Errorf("Expected no error on ExtractResponseBody(), got %v", err)
	}

	if val.Status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, val.Status)
	}

	if val.Message != "Extract Test" {
		t.Errorf("Expected message 'Extract Test', got '%s'", val.Message)
	}

	dataMap, ok := val.Data.(map[string]any)
	if !ok || dataMap["extract"] != "test" {
		t.Errorf("Expected data map with extract 'test', got '%v'", val.Data)
	}
}

func TestWriteJson(t *testing.T) {
	w := httptest.NewRecorder()
	err := writeJSON(w, http.StatusTeapot, true, "I'm a teapot", map[string]string{"tea": "yes"})
	if err != nil {
		t.Errorf("Expected no error on writeJSON(), got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusTeapot {
		t.Errorf("Expected status code %d, got %d", http.StatusTeapot, resp.StatusCode)
	}

	val, err := ExtractResponseBody[NewResponse](resp)
	if err != nil {
		t.Errorf("Expected no error on ExtractResponseBody(), got %v", err)
	}

	if val.Status != http.StatusTeapot {
		t.Errorf("Expected status %d, got %d", http.StatusTeapot, val.Status)
	}

	if val.Message != "I'm a teapot" {
		t.Errorf("Expected message 'I'm a teapot', got '%s'", val.Message)
	}

	dataMap, ok := val.Data.(map[string]any)
	if !ok || dataMap["tea"] != "yes" {
		t.Errorf("Expected data map with tea 'yes', got '%v'", val.Data)
	}
}

func TestWriteJSON_NilWriter(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when http.ResponseWriter is nil, but did not panic")
		}
	}()

	_ = writeJSON(nil, http.StatusOK, true, "This should panic", nil)
}

func TestNewResponseBuilder(t *testing.T) {
	w := httptest.NewRecorder()
	responseBuilder := newResponseBuilder(w, http.StatusOK, false)

	if responseBuilder.Response.Status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, responseBuilder.Response.Status)
	}

	if responseBuilder.Response.Success != true {
		t.Errorf("Expected success to be true, got %v", responseBuilder.Response.Success)
	}

	if responseBuilder.isError != false {
		t.Errorf("Expected isError to be false, got %v", responseBuilder.isError)
	}

	if responseBuilder.Response.Message != "" {
		t.Errorf("Expected empty message, got '%s'", responseBuilder.Response.Message)
	}
}

func TestGetTimestamp(t *testing.T) {
	now := time.Now()
	timestamp := getTimestamp()

	if timestamp.Before(now) || timestamp.After(time.Now()) {
		t.Errorf("Expected timestamp to be around now, got %v", timestamp)
	}

	if !timestamp.Equal(timestamp.UTC()) {
		t.Errorf("Expected timestamp to be in UTC, got %v", timestamp.Location())
	}
}