package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewOK(t *testing.T) {
	w := httptest.NewRecorder()
	err := NewOK(w,
		WithMessage("Test Message"),
		WithData(map[string]string{"key": "value"}),
		WithStatus(http.StatusAccepted),
		Send(),
	)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Expected status %d, got %d", http.StatusAccepted, resp.StatusCode)
	}

	var response NewResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Message() != "Test Message" {
		t.Errorf("Expected message 'Test Message', got '%s'", response.Message())
	}

	dataMap, ok := response.Data().(map[string]any)
	if !ok || dataMap["key"] != "value" {
		t.Errorf("Expected data map with key 'value', got '%v'", response.Data())
	}

	if !response.Success() {
		t.Errorf("Expected success to be true, got false")
	}
}

func TestNewErr(t *testing.T) {
	w := httptest.NewRecorder()
	err := NewErr(w,
		WithMessage("Error Message"),
		WithData("Error Details"),
		WithStatus(http.StatusBadRequest),
		Send(),
	)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	var response NewResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Message() != "Error Message" {
		t.Errorf("Expected message 'Error Message', got '%s'", response.Message())
	}

	// Error responses should include data in the JSON output
	if response.Data() != "Error Details" {
		t.Errorf("Expected data to 'Error Details' for error responses, got '%v'", response.Data())
	}

	if response.Success() {
		t.Errorf("Expected success to be false, got true")
	}
}

func TestErrorWithData(t *testing.T) {
	w := httptest.NewRecorder()
	err := NewErr(w,
		WithStatus(http.StatusBadRequest),
		WithMessage("Bad Request"),
		WithData(map[string]string{"field": "invalid"}),
		Send(),
	)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	var response NewResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Message() != "Bad Request" {
		t.Errorf("Expected message 'Bad Request', got '%s'", response.Message())
	}

	// Error responses should include data
	dataMap, ok := response.Data().(map[string]interface{})
	if !ok || dataMap["field"] != "invalid" {
		t.Errorf("Expected data map with field 'invalid', got '%v'", response.Data())
	}
}

func TestSuccessWithData(t *testing.T) {
	w := httptest.NewRecorder()
	err := NewOK(w,
		WithStatus(http.StatusOK),
		WithMessage("Success"),
		WithData(map[string]string{"result": "ok"}),
		Send(),
	)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var response NewResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Message() != "Success" {
		t.Errorf("Expected message 'Success', got '%s'", response.Message())
	}

	dataMap, ok := response.Data().(map[string]interface{})
	if !ok || dataMap["result"] != "ok" {
		t.Errorf("Expected data map with result 'ok', got '%v'", response.Data())
	}
}

func TestSendError(t *testing.T) {
	w := httptest.NewRecorder()
	err := NewErr(w,
		WithMessage("Error Sending"),
		WithData(map[string]string{"error": "details"}),
		WithStatus(http.StatusInternalServerError),
		Send(),
	)

	if err != nil {
		t.Errorf("Expected no error on Send() for error response, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, resp.StatusCode)
	}

	var response NewResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Error responses should include data
	dataMap, ok := response.Data().(map[string]interface{})
	if !ok || dataMap["error"] != "details" {
		t.Errorf("Expected data map with error 'details', got '%v'", response.Data())
	}
}

func TestWithData(t *testing.T) {
	w := httptest.NewRecorder()
	err := NewOK(w,
		WithData([]int{1, 2, 3}),
		Send(),
	)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	var response NewResponse
	if err := json.NewDecoder(w.Result().Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	dataSlice, ok := response.Data().([]any)
	if !ok || len(dataSlice) != 3 {
		t.Errorf("Expected data slice [1, 2, 3], got '%v'", response.Data())
	}
}

func TestWithMessage(t *testing.T) {
	w := httptest.NewRecorder()
	err := NewOK(w,
		WithMessage("Custom Message"),
		Send(),
	)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	var response NewResponse
	if err := json.NewDecoder(w.Result().Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Message() != "Custom Message" {
		t.Errorf("Expected message 'Custom Message', got '%s'", response.Message())
	}
}

func TestWithStatus(t *testing.T) {
	w := httptest.NewRecorder()
	err := NewOK(w,
		WithStatus(http.StatusCreated),
		Send(),
	)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	var response NewResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if response.Status() != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, response.Status())
	}
}

func TestWithoutSend(t *testing.T) {
	w := httptest.NewRecorder()
	// Without Send(), the response should not be written
	err := NewOK(w,
		WithMessage("Not sent"),
		WithData(map[string]string{"key": "value"}),
	)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Response should be empty since Send() was not called
	if w.Body.Len() != 0 {
		t.Errorf("Expected empty response body, got %d bytes", w.Body.Len())
	}
}
