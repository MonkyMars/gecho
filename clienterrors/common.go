package clienterrors

import (
	"net/http"

	"github.com/MonkyMars/gecho/pkg"
)

func BadRequest(payload *pkg.Payload) error {
	return pkg.Err(payload.W, &pkg.NewResponse{
		Status:  pkg.ValidateStatus(payload.Status, http.StatusBadRequest),
		Message: pkg.ValidateMessage(payload.Message, "Bad Request"),
	})
}

func Unauthorized(payload *pkg.Payload) error {
	return pkg.Err(payload.W, &pkg.NewResponse{
		Status:  pkg.ValidateStatus(payload.Status, http.StatusUnauthorized),
		Message: pkg.ValidateMessage(payload.Message, "Unauthorized"),
	})
}

func Forbidden(payload *pkg.Payload) error {
	return pkg.Err(payload.W, &pkg.NewResponse{
		Status:  http.StatusForbidden,
		Message: pkg.ValidateMessage(payload.Message, "Forbidden"),
	})
}

func NotFound(payload *pkg.Payload) error {
	return pkg.Err(payload.W, &pkg.NewResponse{
		Status:  http.StatusNotFound,
		Message: pkg.ValidateMessage(payload.Message, "Not Found"),
	})
}
