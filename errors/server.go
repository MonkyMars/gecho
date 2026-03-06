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

// NotImplemented sends a 501 Not Implemented response with optional configuration
// Example: errors.NotImplemented(w, gecho.Send())
func NotImplemented(w http.ResponseWriter, opts ...utils.ResponseOption) *utils.Response {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusNotImplemented),
		utils.WithMessage(utils.NotImplementedMessage),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewErr(w, allOpts...)
}

// BadGateway sends a 502 Bad Gateway response with optional configuration
// Example: errors.BadGateway(w, gecho.Send())
func BadGateway(w http.ResponseWriter, opts ...utils.ResponseOption) *utils.Response {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusBadGateway),
		utils.WithMessage(utils.BadGatewayMessage),
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

// GatewayTimeout sends a 504 Gateway Timeout response with optional configuration
// Example: errors.GatewayTimeout(w, gecho.Send())
func GatewayTimeout(w http.ResponseWriter, opts ...utils.ResponseOption) *utils.Response {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusGatewayTimeout),
		utils.WithMessage(utils.GatewayTimeoutMessage),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewErr(w, allOpts...)
}
