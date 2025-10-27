package pkg

import (
	"encoding/json"
	"net/http"
	"time"
)

const SUCCESS = true
const FAILURE = false

// NewResponse is a struct that holds the response data for API responses
type NewResponse struct {
	Status    int       `json:"status"`    // HTTP status code, read https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status for more info
	Success   bool      `json:"success"`   // Indicates whether the request was successful or not
	Message   string    `json:"message"`   // Can be used for both error and success messages
	Data      any       `json:"data"`      // Holds the actual data, returned
	Timestamp time.Time `json:"timestamp"` // Unix timestamp of when the response was generated
}

// getTimestamp returns the current time
func getTimestamp() time.Time {
	return time.Now()
}

// writeJSON writes a JSON response to the http.ResponseWriter
func writeJSON(w http.ResponseWriter, status int, success bool, message string, data any) error {
	if w == nil {
		panic("http.ResponseWriter is nil")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(NewResponse{
		Status:    status,
		Success:   success,
		Message:   message,
		Data:      data,
		Timestamp: getTimestamp(),
	})
}
