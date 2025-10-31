package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewOK(t *testing.T) {
	w := httptest.NewRecorder()
	rb := NewOK(w).
		WithMessage("Test Message").
		WithData(map[string]string{"key": "value"}).
		WithStatus(http.StatusAccepted)

	if rb.Response.Status != http.StatusAccepted {
		t.Errorf("Expected status %d, got %d", http.StatusAccepted, rb.Response.Status)
	}

	if rb.Response.Message != "Test Message" {
		t.Errorf("Expected message 'Test Message', got '%s'", rb.Response.Message)
	}

	dataMap, ok := rb.Response.Data.(map[string]string)
	if !ok || dataMap["key"] != "value" {
		t.Errorf("Expected data map with key 'value', got '%v'", rb.Response.Data)
	}
}

func TestNewErr(t *testing.T) {
	w := httptest.NewRecorder()
	rb := NewErr(w).
		WithMessage("Error Message").
		WithData("Error Details").
		WithStatus(http.StatusBadRequest)

	if rb.Response.Status != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, rb.Response.Status)
	}

	if rb.Response.Message != "Error Message" {
		t.Errorf("Expected message 'Error Message', got '%s'", rb.Response.Message)
	}

	if rb.Response.Data != "Error Details" {
		t.Errorf("Expected data 'Error Details', got '%v'", rb.Response.Data)
	}
}

func TestErrorWithData(t *testing.T) {
	w := httptest.NewRecorder()
	rb := NewErr(w).
		WithStatus(http.StatusBadRequest).
		WithMessage("Bad Request").
		WithData(map[string]string{"field": "invalid"})

	if rb.Response.Status != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, rb.Response.Status)
	}

	if rb.Response.Message != "Bad Request" {
		t.Errorf("Expected message 'Bad Request', got '%s'", rb.Response.Message)
	}

	dataMap, ok := rb.Response.Data.(map[string]string)
	if !ok || dataMap["field"] != "invalid" {
		t.Errorf("Expected data map with field 'invalid', got '%v'", rb.Response.Data)
	}
}

func TestSuccessWithData(t *testing.T) {
	w := httptest.NewRecorder()
	rb := NewOK(w).
		WithStatus(http.StatusOK).
		WithMessage("Success").
		WithData(map[string]string{"result": "ok"})

	if rb.Response.Status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rb.Response.Status)
	}

	if rb.Response.Message != "Success" {
		t.Errorf("Expected message 'Success', got '%s'", rb.Response.Message)
	}

	dataMap, ok := rb.Response.Data.(map[string]string)
	if !ok || dataMap["result"] != "ok" {
		t.Errorf("Expected data map with result 'ok', got '%v'", rb.Response.Data)
	}
}

func TestSendError(t *testing.T) {
	w := httptest.NewRecorder()
	rb := NewErr(w).
		WithMessage("Error Sending").
		WithData(map[string]string{"error": "details"}).
		WithStatus(http.StatusInternalServerError)

	err := rb.Send()
	if err != nil {
		t.Errorf("Expected no error on Send() for error response, got %v", err)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, resp.StatusCode)
	}
}

func TestWithData(t *testing.T) {
	w := httptest.NewRecorder()
	rb := NewOK(w).
		WithData([]int{1, 2, 3})

	dataSlice, ok := rb.Response.Data.([]int)
	if !ok || len(dataSlice) != 3 || dataSlice[0] != 1 || dataSlice[1] != 2 || dataSlice[2] != 3 {
		t.Errorf("Expected data slice [1, 2, 3], got '%v'", rb.Response.Data)
	}
}

func TestWithMessage(t *testing.T) {
	w := httptest.NewRecorder()
	rb := NewOK(w).
		WithMessage("Custom Message")

	if rb.Response.Message != "Custom Message" {
		t.Errorf("Expected message 'Custom Message', got '%s'", rb.Response.Message)
	}
}

func TestWithStatus(t *testing.T) {
	w := httptest.NewRecorder()
	rb := NewOK(w).
		WithStatus(http.StatusCreated)

	if rb.Response.Status != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, rb.Response.Status)
	}
}
