package errors

import (
	"net/http"

	"github.com/MonkyMars/gecho/utils"
)

// BadRequest sends a 400 Bad Request response with optional configuration
// Example: errors.BadRequest(w, gecho.WithData(validationErrors), gecho.Send())
func BadRequest(w http.ResponseWriter, opts ...utils.ResponseOption) error {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusBadRequest),
		utils.WithMessage(utils.BadRequestMessage),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewErr(w, allOpts...)
}

// Unauthorized sends a 401 Unauthorized response with optional configuration
// Example: errors.Unauthorized(w, gecho.Send())
func Unauthorized(w http.ResponseWriter, opts ...utils.ResponseOption) error {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusUnauthorized),
		utils.WithMessage(utils.UnauthorizedMessage),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewErr(w, allOpts...)
}

// Forbidden sends a 403 Forbidden response with optional configuration
// Example: errors.Forbidden(w, gecho.WithMessage("Access denied"), gecho.Send())
func Forbidden(w http.ResponseWriter, opts ...utils.ResponseOption) error {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusForbidden),
		utils.WithMessage(utils.ForbiddenMessage),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewErr(w, allOpts...)
}

// NotFound sends a 404 Not Found response with optional configuration
// Example: errors.NotFound(w, gecho.Send())
func NotFound(w http.ResponseWriter, opts ...utils.ResponseOption) error {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusNotFound),
		utils.WithMessage(utils.NotFoundMessage),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewErr(w, allOpts...)
}

// MethodNotAllowed sends a 405 Method Not Allowed response with optional configuration
// Example: errors.MethodNotAllowed(w, gecho.Send())
func MethodNotAllowed(w http.ResponseWriter, opts ...utils.ResponseOption) error {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusMethodNotAllowed),
		utils.WithMessage(utils.MethodNotAllowedMessage),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewErr(w, allOpts...)
}

// Conflict sends a 409 Conflict response with optional configuration
// Example: errors.Conflict(w, gecho.WithMessage("Resource already exists"), gecho.Send())
func Conflict(w http.ResponseWriter, opts ...utils.ResponseOption) error {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusConflict),
		utils.WithMessage(utils.ConflictMessage),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewErr(w, allOpts...)
}
