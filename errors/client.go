package errors

import (
	"net/http"

	"github.com/MonkyMars/gecho/utils"
)

// BadRequest returns a ResponseBuilder for 400 Bad Request responses
// Use Send() to send the response with the default values
func BadRequest(w http.ResponseWriter) *utils.ResponseBuilder {
	return utils.NewErr(w).WithStatus(http.StatusBadRequest).
		WithMessage("Bad Request")
}

// Unauthorized returns a ResponseBuilder for 401 Unauthorized responses
// Use Send() to send the response with the default values
func Unauthorized(w http.ResponseWriter) *utils.ResponseBuilder {
	return utils.NewErr(w).WithStatus(http.StatusUnauthorized).
		WithMessage("Unauthorized")
}

// Forbidden returns a ResponseBuilder for 403 Forbidden responses
// Use Send() to send the response with the default values
func Forbidden(w http.ResponseWriter) *utils.ResponseBuilder {
	return utils.NewErr(w).WithStatus(http.StatusForbidden).
		WithMessage("Forbidden")
}

// NotFound returns a ResponseBuilder for 404 Not Found responses
// Use Send() to send the response with the default values
func NotFound(w http.ResponseWriter) *utils.ResponseBuilder {
	return utils.NewErr(w).WithStatus(http.StatusNotFound).
		WithMessage("Not Found")
}

// MethodNotAllowed returns a ResponseBuilder for 405 Method Not Allowed responses
// Use Send() to send the response with the default values
func MethodNotAllowed(w http.ResponseWriter) *utils.ResponseBuilder {
	return utils.NewErr(w).WithStatus(http.StatusMethodNotAllowed).
		WithMessage("Method Not Allowed")
}
