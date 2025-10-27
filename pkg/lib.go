package pkg

import (
	"encoding/json"
	"net/http"
	"time"
)

// NewResponse is a struct that holds the response data for API responses
type NewResponse struct {
	Status    int       `json:"status"`    // HTTP status code, read https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status for more info
	Message   string    `json:"message"`   // Can be used for both error and success messages
	Data      any       `json:"data"`      // Holds the actual data, returned
	Timestamp time.Time `json:"timestamp"` // Unix timestamp of when the response was generated
}

type Payload struct {
	W       http.ResponseWriter
	Status  int
	Message string
	Data    any
}

func GetTimestamp() time.Time {
	return time.Now()
}

func WriteJSON(w http.ResponseWriter, status int, message string, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(NewResponse{
		Status:    status,
		Message:   message,
		Data:      data,
		Timestamp: GetTimestamp(),
	})
}

func ValidateMessage(message, alternative string) string {
	if message == "" {
		return alternative
	}
	return message
}

func ValidateStatus(status, alternative int) int {
	if status == 0 {
		return alternative
	}
	return status
}
