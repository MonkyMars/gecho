package utils

import (
	"net/http"
	"runtime"
	"sync"
)

// ResponseBuilder provides a fluent interface for building responses
type ResponseBuilder struct {
	w        http.ResponseWriter
	response *NewResponse
	isError  bool
	sent     bool
	mu       sync.Mutex
}

// NewOK creates a new ResponseBuilder for success responses
func NewOK(w http.ResponseWriter) *ResponseBuilder {
	rb := newResponseBuilder(w, http.StatusOK, false)
	return rb
}

// NewErr creates a new ResponseBuilder for error responses
func NewErr(w http.ResponseWriter) *ResponseBuilder {
	rb := newResponseBuilder(w, http.StatusInternalServerError, true)
	return rb
}

// NewResponseBuilder creates a ResponseBuilder with custom status and type
func newResponseBuilder(w http.ResponseWriter, status int, isError bool) *ResponseBuilder {
	success := SUCCESS
	if isError {
		success = FAILURE
	}

	rb := &ResponseBuilder{
		w: w,
		response: &NewResponse{
			Status:    status,
			Success:   success,
			Timestamp: getTimestamp(),
		},
		isError: isError,
	}
	runtime.SetFinalizer(rb, (*ResponseBuilder).finalize)
	return rb
}

// WithMessage sets the response message and auto-sends
func (rb *ResponseBuilder) WithMessage(message string) *ResponseBuilder {
	rb.response.Message = message
	return rb
}

// WithData sets the response data and returns builder for chaining
func (rb *ResponseBuilder) WithData(data any) *ResponseBuilder {
	if !rb.isError {
		rb.response.Data = data
	}
	return rb
}

// WithStatus sets the HTTP status code and returns builder for chaining
func (rb *ResponseBuilder) WithStatus(status int) *ResponseBuilder {
	rb.response.Status = status
	return rb
}

// Send manually sends the response
func (rb *ResponseBuilder) Send() error {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	if rb.sent {
		return nil
	}

	data := rb.response.Data
	if rb.isError {
		data = nil
	}

	err := writeJSON(rb.w, rb.response.Status, rb.response.Success, rb.response.Message, data)
	rb.sent = true
	runtime.SetFinalizer(rb, nil) // Clear finalizer since we've sent
	return err
}

func (rb *ResponseBuilder) finalize() {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	if !rb.sent {
		data := rb.response.Data
		if rb.isError {
			data = nil
		}
		_ = writeJSON(rb.w, rb.response.Status, rb.response.Success, rb.response.Message, data)
		rb.sent = true
	}
}
