package servererrors

import (
	"net/http"

	"github.com/MonkyMars/gecho/pkg"
)

// InternalServerError sends a 500 Internal Server Error response with an optional custom message.
func InternalServerError(payload *pkg.Payload) error {
	return pkg.Err(payload.W, &pkg.NewResponse{
		Status:  pkg.ValidateStatus(payload.Status, http.StatusInternalServerError),
		Message: pkg.ValidateMessage(payload.Message, "Internal Server Error"),
	})
}

// ServiceUnavailable sends a 503 Service Unavailable response with an optional custom message.
func ServiceUnavailable(payload *pkg.Payload) error {
	return pkg.Err(payload.W, &pkg.NewResponse{
		Status:  pkg.ValidateStatus(payload.Status, http.StatusServiceUnavailable),
		Message: pkg.ValidateMessage(payload.Message, "Service Unavailable"),
	})
}
