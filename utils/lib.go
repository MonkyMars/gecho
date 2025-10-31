package utils

import (
	"encoding/json"
	"net/http"
	"time"
)


const NOT_FOUND_MESSAGE = "Resource not found"
const INTERNAL_SERVER_ERROR_MESSAGE = "Internal server error"
const SERVICE_UNAVAILABLE_MESSAGE = "Service unavailable"
const BAD_REQUEST_MESSAGE = "Bad request"
const UNAUTHORIZED_MESSAGE = "Unauthorized"
const FORBIDDEN_MESSAGE = "Forbidden"
const METHOD_NOT_ALLOWED_MESSAGE = "Method not allowed"

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

func ExtractResponseBody[T any](resp *http.Response) (T, error) {
	var result T
	err := json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}