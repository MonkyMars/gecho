package errors

import (
	"net/http"

	"github.com/MonkyMars/gecho/utils"
)

// BadRequest returns a ResponseBuilder for 400 Bad Request responses
// Use Send() to send the response with the default values
// You can use WithData to add more details about the bad request, such as validation errors
func BadRequest(w http.ResponseWriter) *utils.ResponseBuilder {
	return utils.NewErr(w).WithStatus(http.StatusBadRequest).
		WithMessage(utils.BadRequestMessage)
}

// Unauthorized returns a ResponseBuilder for 401 Unauthorized responses
// Use Send() to send the response with the default values
func Unauthorized(w http.ResponseWriter) *utils.ResponseBuilder {
	return utils.NewErr(w).WithStatus(http.StatusUnauthorized).
		WithMessage(utils.UnauthorizedMessage)
}

// Forbidden returns a ResponseBuilder for 403 Forbidden responses
// Use Send() to send the response with the default values
func Forbidden(w http.ResponseWriter) *utils.ResponseBuilder {
	return utils.NewErr(w).WithStatus(http.StatusForbidden).
		WithMessage(utils.ForbiddenMessage)
}

// NotFound returns a ResponseBuilder for 404 Not Found responses
// Use Send() to send the response with the default values
func NotFound(w http.ResponseWriter) *utils.ResponseBuilder {
	return utils.NewErr(w).WithStatus(http.StatusNotFound).
		WithMessage(utils.NotFoundMessage)
}

// MethodNotAllowed returns a ResponseBuilder for 405 Method Not Allowed responses
// Use Send() to send the response with the default values
func MethodNotAllowed(w http.ResponseWriter) *utils.ResponseBuilder {
	return utils.NewErr(w).WithStatus(http.StatusMethodNotAllowed).
		WithMessage(utils.MethodNotAllowedMessage)
}

// Conflict returns a ResponseBuilder for 409 Conflict responses
// Use Send() to send the response with the default values
func Conflict(w http.ResponseWriter) *utils.ResponseBuilder {
	return utils.NewErr(w).WithStatus(http.StatusConflict).
		WithMessage(utils.ConflictMessage)
}
