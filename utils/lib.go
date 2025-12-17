package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

const NotFoundMessage = "Resource not found"
const InternalServerErrorMessage = "Internal server error"
const ServiceUnavailableMessage = "Service unavailable"
const BadRequestMessage = "Bad request"
const UnauthorizedMessage = "Unauthorized"
const ForbiddenMessage = "Forbidden"
const MethodNotAllowedMessage = "Method not allowed"
const ConflictMessage = "Conflict"
const TooManyRequestsMessage = "Too many requests"

// NewResponse is a struct that holds the response data for API responses
type NewResponse struct {
	status    int       // HTTP status code, read https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status for more info
	success   bool      // Indicates whether the request was successful or not
	message   string    // Can be used for both error and success messages
	data      any       // Holds the actual data, returned
	timestamp time.Time // Unix timestamp of when the response was generated
}

func (nr *NewResponse) Status() int {
	return nr.status
}

func (nr *NewResponse) Message() string {
	return nr.message
}

func (nr *NewResponse) Data() any {
	return nr.data
}

func (nr *NewResponse) Timestamp() time.Time {
	return nr.timestamp
}

func (nr *NewResponse) Success() bool {
	return nr.success
}

// getTimestamp returns the current time
func getTimestamp() time.Time {
	return time.Now()
}

// MarshalJSON implements custom marshaling for NewResponse so we can keep
// fields unexported but still produce JSON with exported keys.
func (nr NewResponse) MarshalJSON() ([]byte, error) {
	type jsonResp struct {
		Status    int       `json:"status"`
		Success   bool      `json:"success"`
		Message   string    `json:"message"`
		Data      any       `json:"data"`
		Timestamp time.Time `json:"timestamp"`
	}

	jr := jsonResp{
		Status:    nr.status,
		Success:   nr.success,
		Message:   nr.message,
		Data:      nr.data,
		Timestamp: nr.timestamp,
	}

	return json.Marshal(jr)
}

// UnmarshalJSON implements custom unmarshaling into the unexported fields of NewResponse.
func (nr *NewResponse) UnmarshalJSON(b []byte) error {
	type jsonResp struct {
		Status    int       `json:"status"`
		Success   bool      `json:"success"`
		Message   string    `json:"message"`
		Data      any       `json:"data"`
		Timestamp time.Time `json:"timestamp"`
	}

	var jr jsonResp
	if err := json.Unmarshal(b, &jr); err != nil {
		return err
	}

	nr.status = jr.Status
	nr.success = jr.Success
	nr.message = jr.Message
	nr.data = jr.Data
	nr.timestamp = jr.Timestamp

	return nil
}

// writeJSON writes a JSON response to the http.ResponseWriter
func writeJSON(w http.ResponseWriter, status int, success bool, message string, data any) error {
	if w == nil {
		panic("http.ResponseWriter is nil")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(NewResponse{
		status:    status,
		success:   success,
		message:   message,
		data:      data,
		timestamp: getTimestamp(),
	})
}

func ExtractResponseBody[T any](resp *http.Response) (T, error) {
	var result T
	err := json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}
