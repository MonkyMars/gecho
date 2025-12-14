package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestResponseModification demonstrates the new Response modification feature
func TestResponseModification(t *testing.T) {
	t.Run("SetMessage", func(t *testing.T) {
		w := httptest.NewRecorder()
		resp := NewOK(w, WithMessage("Original message"))

		// Modify message before sending
		resp.SetMessage("Modified message")
		resp.Send()

		var response NewResponse
		if err := json.NewDecoder(w.Result().Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if response.Message() != "Modified message" {
			t.Errorf("Expected message 'Modified message', got '%s'", response.Message())
		}
	})

	t.Run("SetStatus", func(t *testing.T) {
		w := httptest.NewRecorder()
		resp := NewOK(w, WithStatus(http.StatusOK))

		// Modify status before sending
		resp.SetStatus(http.StatusCreated)
		resp.Send()

		if w.Result().StatusCode != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Result().StatusCode)
		}
	})

	t.Run("SetData", func(t *testing.T) {
		w := httptest.NewRecorder()
		resp := NewOK(w, WithData(map[string]string{"initial": "data"}))

		// Replace entire data
		resp.SetData(map[string]string{"replaced": "data"})
		resp.Send()

		var response NewResponse
		if err := json.NewDecoder(w.Result().Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		dataMap, ok := response.Data().(map[string]any)
		if !ok {
			t.Fatalf("Expected data to be a map")
		}

		if dataMap["replaced"] != "data" {
			t.Errorf("Expected replaced='data', got %v", dataMap["replaced"])
		}

		if _, exists := dataMap["initial"]; exists {
			t.Errorf("Expected 'initial' key to be removed")
		}
	})

	t.Run("AddData", func(t *testing.T) {
		w := httptest.NewRecorder()
		resp := NewOK(w, WithData(map[string]string{"initial": "value"}))

		// Add additional fields
		resp.AddData("added", "field")
		resp.AddData("another", 123)
		resp.Send()

		var response NewResponse
		if err := json.NewDecoder(w.Result().Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		dataMap, ok := response.Data().(map[string]any)
		if !ok {
			t.Fatalf("Expected data to be a map")
		}

		if dataMap["initial"] != "value" {
			t.Errorf("Expected initial='value', got %v", dataMap["initial"])
		}

		if dataMap["added"] != "field" {
			t.Errorf("Expected added='field', got %v", dataMap["added"])
		}

		if dataMap["another"] != float64(123) { // JSON numbers are float64
			t.Errorf("Expected another=123, got %v", dataMap["another"])
		}
	})

	t.Run("ChainedModifications", func(t *testing.T) {
		w := httptest.NewRecorder()
		resp := NewOK(w)

		// Chain multiple modifications
		resp.
			SetMessage("Chained message").
			SetStatus(http.StatusAccepted).
			AddData("key1", "value1").
			AddData("key2", "value2").
			Send()

		if w.Result().StatusCode != http.StatusAccepted {
			t.Errorf("Expected status code %d, got %d", http.StatusAccepted, w.Result().StatusCode)
		}

		var response NewResponse
		if err := json.NewDecoder(w.Result().Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if response.Message() != "Chained message" {
			t.Errorf("Expected message 'Chained message', got '%s'", response.Message())
		}

		dataMap, ok := response.Data().(map[string]any)
		if !ok {
			t.Fatalf("Expected data to be a map")
		}

		if dataMap["key1"] != "value1" || dataMap["key2"] != "value2" {
			t.Errorf("Expected key1='value1' and key2='value2', got %v", dataMap)
		}
	})

	t.Run("ConditionalModification", func(t *testing.T) {
		w := httptest.NewRecorder()
		resp := NewOK(w, WithData(map[string]string{"user": "alice"}))

		// Simulate conditional logic
		isAdmin := true
		if isAdmin {
			resp.AddData("role", "admin")
			resp.AddData("permissions", []string{"read", "write", "delete"})
		}

		resp.Send()

		var response NewResponse
		if err := json.NewDecoder(w.Result().Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		dataMap, ok := response.Data().(map[string]any)
		if !ok {
			t.Fatalf("Expected data to be a map")
		}

		if dataMap["role"] != "admin" {
			t.Errorf("Expected role='admin', got %v", dataMap["role"])
		}

		if dataMap["user"] != "alice" {
			t.Errorf("Expected user='alice', got %v", dataMap["user"])
		}
	})

	t.Run("ModifyAfterCreation", func(t *testing.T) {
		w := httptest.NewRecorder()

		// Create response without any options
		resp := NewOK(w)

		// Build response step by step
		resp.SetMessage("Step by step")
		resp.SetStatus(http.StatusOK)
		resp.AddData("step1", "complete")
		resp.AddData("step2", "complete")
		resp.AddData("step3", "complete")

		resp.Send()

		var response NewResponse
		if err := json.NewDecoder(w.Result().Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if response.Message() != "Step by step" {
			t.Errorf("Expected message 'Step by step', got '%s'", response.Message())
		}

		dataMap, ok := response.Data().(map[string]any)
		if !ok {
			t.Fatalf("Expected data to be a map")
		}

		if len(dataMap) != 3 {
			t.Errorf("Expected 3 data fields, got %d", len(dataMap))
		}
	})

	t.Run("AddDataToNonMapData", func(t *testing.T) {
		w := httptest.NewRecorder()
		resp := NewOK(w, WithData("simple string"))

		// AddData should convert the data to a map
		resp.AddData("extra", "value")
		resp.Send()

		var response NewResponse
		if err := json.NewDecoder(w.Result().Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		dataMap, ok := response.Data().(map[string]any)
		if !ok {
			t.Fatalf("Expected data to be converted to a map")
		}

		if dataMap["data"] != "simple string" {
			t.Errorf("Expected data='simple string', got %v", dataMap["data"])
		}

		if dataMap["extra"] != "value" {
			t.Errorf("Expected extra='value', got %v", dataMap["extra"])
		}
	})

	t.Run("NilResponseSafety", func(t *testing.T) {
		var resp *Response = nil

		// All methods should handle nil safely
		resp = resp.SetMessage("test")
		resp = resp.SetStatus(200)
		resp = resp.SetData("test")
		resp = resp.AddData("key", "value")
		err := resp.Send()

		if resp != nil {
			t.Errorf("Expected nil response to remain nil")
		}

		if err != nil {
			t.Errorf("Expected no error from nil response, got %v", err)
		}
	})

	t.Run("BackwardCompatibility", func(t *testing.T) {
		w := httptest.NewRecorder()

		// Old pattern with Send() should still work and return nil
		resp := NewOK(w, WithMessage("Immediate send"), Send())

		if resp != nil {
			t.Errorf("Expected nil response when using Send() option, got %v", resp)
		}

		var response NewResponse
		if err := json.NewDecoder(w.Result().Body).Decode(&response); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if response.Message() != "Immediate send" {
			t.Errorf("Expected message 'Immediate send', got '%s'", response.Message())
		}
	})
}
