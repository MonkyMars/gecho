package success

import (
	"net/http"

	"github.com/MonkyMars/gecho/utils"
)

// Success sends a 200 OK response with optional configuration
// Example: success.Success(w, gecho.WithData(userData), gecho.Send())
func Success(w http.ResponseWriter, opts ...utils.ResponseOption) error {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusOK),
		utils.WithMessage("Success"),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewOK(w, allOpts...)
}

// Created sends a 201 Created response with optional configuration
// Example: success.Created(w, gecho.WithData(newResource), gecho.Send())
func Created(w http.ResponseWriter, opts ...utils.ResponseOption) error {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusCreated),
		utils.WithMessage("Resource Created"),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewOK(w, allOpts...)
}

// Accepted sends a 202 Accepted response with optional configuration
// Example: success.Accepted(w, gecho.WithMessage("Request accepted for processing"), gecho.Send())
func Accepted(w http.ResponseWriter, opts ...utils.ResponseOption) error {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusAccepted),
		utils.WithMessage("Accepted"),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewOK(w, allOpts...)
}

// NoContent sends a 204 No Content response with optional configuration
// Example: success.NoContent(w, gecho.Send())
func NoContent(w http.ResponseWriter, opts ...utils.ResponseOption) error {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusNoContent),
		utils.WithMessage("No Content"),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewOK(w, allOpts...)
}
