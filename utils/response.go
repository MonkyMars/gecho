package utils

import (
	"net/http"
)

// Response represents an HTTP response that can be modified before sending
type Response struct {
	w       http.ResponseWriter
	status  int
	success bool
	message string
	data    any
}

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

// SetMessage sets the response message
func (r *Response) SetMessage(message string) *Response {
	if r == nil {
		return nil
	}
	r.message = message
	return r
}

// SetStatus sets the HTTP status code
func (r *Response) SetStatus(status int) *Response {
	if r == nil {
		return nil
	}
	r.status = status
	return r
}

// SetData replaces the entire response data
func (r *Response) SetData(data any) *Response {
	if r == nil {
		return nil
	}
	r.data = data
	return r
}

// AddData adds a single key-value pair to the response data
// If data is not a map, it will be converted to one
func (r *Response) AddData(key string, value any) *Response {
	if r == nil {
		return nil
	}

	// Ensure data is a map
	if r.data == nil {
		r.data = make(map[string]any)
	} else if dataMap, ok := r.data.(map[string]any); ok {
		// Already map[string]any, use it directly
		dataMap[key] = value
		return r
	} else {
		// Try to convert from other map types or wrap non-map data
		newMap := make(map[string]any)

		// Try reflection to convert from other map types (e.g., map[string]string)
		switch v := r.data.(type) {
		case map[string]string:
			for k, val := range v {
				newMap[k] = val
			}
		case map[string]int:
			for k, val := range v {
				newMap[k] = val
			}
		default:
			// Wrap non-map data in a "data" field
			newMap["data"] = r.data
		}

		r.data = newMap
	}

	r.data.(map[string]any)[key] = value
	return r
}

// Send writes the response to the HTTP response writer
func (r *Response) Send() error {
	if r == nil {
		return nil
	}

	return writeJSON(r.w, r.status, r.success, r.message, r.data)
}

// buildResponse applies all options and returns a Response object
func buildResponse(w http.ResponseWriter, defaultStatus int, isError bool, defaultMessage string, opts []ResponseOption) *Response {
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

	// Create Response object
	resp := &Response{
		w:       config.w,
		status:  config.status,
		success: config.success,
		message: config.message,
		data:    nil,
	}

	// Set data as-is
	resp.data = config.data

	// Auto-send if Send() option was provided
	if config.send {
		resp.Send()
		return nil
	}

	return resp
}

// NewOK creates a success response with options
func NewOK(w http.ResponseWriter, opts ...ResponseOption) *Response {
	return buildResponse(w, http.StatusOK, false, "Success", opts)
}

// NewErr creates an error response with options
func NewErr(w http.ResponseWriter, opts ...ResponseOption) *Response {
	return buildResponse(w, http.StatusInternalServerError, true, "Internal server error", opts)
}
