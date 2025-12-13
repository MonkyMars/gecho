package utils

import (
	"net/http"
)

// ResponseOption is a function that configures a response
type ResponseOption func(*responseConfig)

// responseConfig holds the configuration for a response
type responseConfig struct {
	w       http.ResponseWriter
	status  int
	success bool
	message string
	data    any
	send    bool
}

// WithData sets the response data
func WithData(data any) ResponseOption {
	return func(rc *responseConfig) {
		rc.data = data
	}
}

// WithMessage sets the response message
func WithMessage(message string) ResponseOption {
	return func(rc *responseConfig) {
		rc.message = message
	}
}

// WithStatus sets the HTTP status code
func WithStatus(status int) ResponseOption {
	return func(rc *responseConfig) {
		rc.status = status
	}
}

// Send marks the response to be sent immediately
func Send() ResponseOption {
	return func(rc *responseConfig) {
		rc.send = true
	}
}

// buildResponse applies all options and optionally sends the response
func buildResponse(w http.ResponseWriter, defaultStatus int, isError bool, defaultMessage string, opts []ResponseOption) error {
	success := !isError

	config := &responseConfig{
		w:       w,
		status:  defaultStatus,
		success: success,
		message: defaultMessage,
		data:    nil,
		send:    false,
	}

	// Apply all options
	for _, opt := range opts {
		opt(config)
	}

	// Send if requested
	if config.send {
		data := config.data
		return writeJSON(config.w, config.status, config.success, config.message, data)
	}

	return nil
}

// NewOK creates a success response with options
func NewOK(w http.ResponseWriter, opts ...ResponseOption) error {
	return buildResponse(w, http.StatusOK, false, "Success", opts)
}

// NewErr creates an error response with options
func NewErr(w http.ResponseWriter, opts ...ResponseOption) error {
	return buildResponse(w, http.StatusInternalServerError, true, "Internal server error", opts)
}
