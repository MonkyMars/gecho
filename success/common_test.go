package success

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MonkyMars/gecho/utils"
)

func TestCorrectSuccessResponses(t *testing.T) {
	tests := []struct {
		name           string
		fn             func(http.ResponseWriter, ...utils.ResponseOption) *utils.Response
		expectedStatus int
		expectedMsg    string
	}{
		{
			name:           "Success",
			fn:             Success,
			expectedStatus: http.StatusOK,
			expectedMsg:    "Success",
		},
		{
			name:           "Created",
			fn:             Created,
			expectedStatus: http.StatusCreated,
			expectedMsg:    "Resource Created",
		},
		{
			name:           "Accepted",
			fn:             Accepted,
			expectedStatus: http.StatusAccepted,
			expectedMsg:    "Accepted",
		},
		{
			name:           "NoContent",
			fn:             NoContent,
			expectedStatus: http.StatusNoContent,
			expectedMsg:    "No Content",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			tt.fn(w, utils.Send())

			resp := w.Result()
			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, resp.StatusCode)
			}

			var response utils.NewResponse
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				t.Fatalf("Failed to decode response: %v", err)
			}

			if response.Message() != tt.expectedMsg {
				t.Errorf("Expected message '%s', got '%s'", tt.expectedMsg, response.Message())
			}

			if !response.Success() {
				t.Errorf("Expected success to be true, got false")
			}
		})
	}
}
