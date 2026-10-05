package errors

import (
	"net/http"

	"github.com/MonkyMars/gecho/utils"
)

// InternalServerError creates a 500 Internal Server Error response with optional configuration.
func InternalServerError(w http.ResponseWriter, opts ...utils.ResponseOption) *utils.Response {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusInternalServerError),
		utils.WithMessage(utils.InternalServerErrorMessage),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewErr(w, allOpts...)
}

// ServiceUnavailable creates a 503 Service Unavailable response with optional configuration.
func ServiceUnavailable(w http.ResponseWriter, opts ...utils.ResponseOption) *utils.Response {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusServiceUnavailable),
		utils.WithMessage(utils.ServiceUnavailableMessage),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewErr(w, allOpts...)
}
