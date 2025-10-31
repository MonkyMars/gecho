package success

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MonkyMars/gecho/utils"
)

func TestCorrectSuccessResponses(t *testing.T) {
	// Test Success response
	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		responseBuilder := Success(w)
		err := responseBuilder.Send()
		if err != nil {
			t.Errorf("Expected no error on Send(), got %v", err)
		}
		if w.Result().StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Result().StatusCode)
		}

		// Extract and verify response body
		val, err := utils.ExtractResponseBody[utils.NewResponse](w.Result())
		if err != nil {
			t.Errorf("Expected no error on ExtractResponseBody(), got %v", err)
		}
		if val.Status() != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, val.Status())
		}
		if val.Message() != "Success" {
			t.Errorf("Expected message 'Success', got '%s'", val.Message())
		}
		if val.Data() != nil {
			t.Errorf("Expected nil data, got '%v'", val.Data())
		}
	})

	// Test Created response
	t.Run("Created", func(t *testing.T) {
		w := httptest.NewRecorder()
		responseBuilder := Created(w)
		err := responseBuilder.Send()
		if err != nil {
			t.Errorf("Expected no error on Send(), got %v", err)
		}
		if w.Result().StatusCode != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Result().StatusCode)
		}

		// Extract and verify response body
		val, err := utils.ExtractResponseBody[utils.NewResponse](w.Result())
		if err != nil {
			t.Errorf("Expected no error on ExtractResponseBody(), got %v", err)
		}
		if val.Status() != http.StatusCreated {
			t.Errorf("Expected status %d, got %d", http.StatusCreated, val.Status())
		}
		if val.Message() != "Resource Created" {
			t.Errorf("Expected message 'Resource Created', got '%s'", val.Message())
		}
	})

	// Test Accepted response
	t.Run("Accepted", func(t *testing.T) {
		w := httptest.NewRecorder()
		responseBuilder := Accepted(w)
		err := responseBuilder.Send()
		if err != nil {
			t.Errorf("Expected no error on Send(), got %v", err)
		}
		if w.Result().StatusCode != http.StatusAccepted {
			t.Errorf("Expected status code %d, got %d", http.StatusAccepted, w.Result().StatusCode)
		}

		// Extract and verify response body
		val, err := utils.ExtractResponseBody[utils.NewResponse](w.Result())
		if err != nil {
			t.Errorf("Expected no error on ExtractResponseBody(), got %v", err)
		}
		if val.Status() != http.StatusAccepted {
			t.Errorf("Expected status %d, got %d", http.StatusAccepted, val.Status())
		}
		if val.Message() != "Accepted" {
			t.Errorf("Expected message 'Accepted', got '%s'", val.Message())
		}
		if val.Data() != nil {
			t.Errorf("Expected nil data, got '%v'", val.Data())
		}
	})

	// Test NoContent response
	t.Run("NoContent", func(t *testing.T) {
		w := httptest.NewRecorder()
		responseBuilder := NoContent(w)
		err := responseBuilder.Send()
		if err != nil {
			t.Errorf("Expected no error on Send(), got %v", err)
		}
		if w.Result().StatusCode != http.StatusNoContent {
			t.Errorf("Expected status code %d, got %d", http.StatusNoContent, w.Result().StatusCode)
		}

		// Extract and verify response body
		val, err := utils.ExtractResponseBody[utils.NewResponse](w.Result())
		if err != nil {
			t.Errorf("Expected no error on ExtractResponseBody(), got %v", err)
		}
		if val.Status() != http.StatusNoContent {
			t.Errorf("Expected status %d, got %d", http.StatusNoContent, val.Status())
		}
		if val.Message() != "No Content" {
			t.Errorf("Expected message 'No Content', got '%s'", val.Message())
		}
		if val.Data() != nil {
			t.Errorf("Expected nil data, got '%v'", val.Data())
		}
	})
}
