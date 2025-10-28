package errors

import (
	"net/http"

	"github.com/MonkyMars/gecho/utils"
)

// InternalServerError returns a ResponseBuilder for 500 Internal Server Error responses
// Use Send() to send the response with the default values
func InternalServerError(w http.ResponseWriter) *utils.ResponseBuilder {
	return utils.NewErr(w).WithStatus(http.StatusInternalServerError).
		WithMessage("Internal Server Error")
}

// ServiceUnavailable returns a ResponseBuilder for 503 Service Unavailable responses
// Use Send() to send the response with the default values
func ServiceUnavailable(w http.ResponseWriter) *utils.ResponseBuilder {
	return utils.NewErr(w).WithStatus(http.StatusServiceUnavailable).
		WithMessage("Service Unavailable")
}
