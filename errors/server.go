package errors

import (
	"net/http"

	"github.com/MonkyMars/gecho/utils"
)

// InternalServerError sends a 500 Internal Server Error response with optional configuration
// Example: errors.InternalServerError(w, gecho.Send())
func InternalServerError(w http.ResponseWriter, opts ...utils.ResponseOption) *utils.Response {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusInternalServerError),
		utils.WithMessage(utils.InternalServerErrorMessage),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewErr(w, allOpts...)
}

// ServiceUnavailable sends a 503 Service Unavailable response with optional configuration
// Example: errors.ServiceUnavailable(w, gecho.WithMessage("Maintenance mode"), gecho.Send())
func ServiceUnavailable(w http.ResponseWriter, opts ...utils.ResponseOption) *utils.Response {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusServiceUnavailable),
		utils.WithMessage(utils.ServiceUnavailableMessage),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewErr(w, allOpts...)
}
