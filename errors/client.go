package errors

import (
	"net/http"

	"github.com/MonkyMars/gecho/utils"
)

// BadRequest creates a 400 Bad Request response with optional configuration.
func BadRequest(w http.ResponseWriter, opts ...utils.ResponseOption) *utils.Response {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusBadRequest),
		utils.WithMessage(utils.BadRequestMessage),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewErr(w, allOpts...)
}

// Unauthorized creates a 401 Unauthorized response with optional configuration.
func Unauthorized(w http.ResponseWriter, opts ...utils.ResponseOption) *utils.Response {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusUnauthorized),
		utils.WithMessage(utils.UnauthorizedMessage),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewErr(w, allOpts...)
}

// Forbidden creates a 403 Forbidden response with optional configuration.
func Forbidden(w http.ResponseWriter, opts ...utils.ResponseOption) *utils.Response {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusForbidden),
		utils.WithMessage(utils.ForbiddenMessage),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewErr(w, allOpts...)
}

// NotFound creates a 404 Not Found response with optional configuration.
func NotFound(w http.ResponseWriter, opts ...utils.ResponseOption) *utils.Response {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusNotFound),
		utils.WithMessage(utils.NotFoundMessage),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewErr(w, allOpts...)
}

// MethodNotAllowed creates a 405 Method Not Allowed response with optional configuration.
func MethodNotAllowed(w http.ResponseWriter, opts ...utils.ResponseOption) *utils.Response {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusMethodNotAllowed),
		utils.WithMessage(utils.MethodNotAllowedMessage),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewErr(w, allOpts...)
}

// Conflict creates a 409 Conflict response with optional configuration.
func Conflict(w http.ResponseWriter, opts ...utils.ResponseOption) *utils.Response {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusConflict),
		utils.WithMessage(utils.ConflictMessage),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewErr(w, allOpts...)
}

// TooManyRequests creates a 429 Too Many Requests response with optional configuration.
func TooManyRequests(w http.ResponseWriter, opts ...utils.ResponseOption) *utils.Response {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusTooManyRequests),
		utils.WithMessage(utils.TooManyRequestsMessage),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewErr(w, allOpts...)
}
