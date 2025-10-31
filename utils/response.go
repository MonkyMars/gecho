package utils

import (
	"net/http"
)

// ResponseBuilder provides a fluent interface for building responses
type ResponseBuilder struct {
	w        http.ResponseWriter
	response *NewResponse
	isError  bool
}

func (rb *ResponseBuilder) Response() *NewResponse {
	return rb.response
}

// NewOK creates a new ResponseBuilder for success responses
func NewOK(w http.ResponseWriter) *ResponseBuilder {
	return newResponseBuilder(w, http.StatusOK, false)
}

// NewErr creates a new ResponseBuilder for error responses
func NewErr(w http.ResponseWriter) *ResponseBuilder {
	return newResponseBuilder(w, http.StatusInternalServerError, true)
}

// newResponseBuilder creates a ResponseBuilder with custom status and type
func newResponseBuilder(w http.ResponseWriter, status int, isError bool) *ResponseBuilder {
	success := true
	if isError {
		success = false
	}

	return &ResponseBuilder{
		w: w,
		response: &NewResponse{
			status:    status,
			success:   success,
			timestamp: getTimestamp(),
		},
		isError: isError,
	}
}

// WithMessage sets the response message and returns builder for chaining
func (rb *ResponseBuilder) WithMessage(message string) *ResponseBuilder {
	rb.Response().message = message
	return rb
}

// WithData sets the response data and returns builder for chaining
func (rb *ResponseBuilder) WithData(data any) *ResponseBuilder {
	rb.Response().data = data
	return rb
}

// WithStatus sets the HTTP status code and returns builder for chaining
func (rb *ResponseBuilder) WithStatus(status int) *ResponseBuilder {
	rb.Response().status = status
	return rb
}

// Send manually sends the response
func (rb *ResponseBuilder) Send() error {
	data := rb.Response().Data()

	if rb.isError {
		data = nil
	}

	err := writeJSON(rb.w, rb.Response().Status(), rb.Response().Success(), rb.Response().Message(), data)
	return err
}
