package success

import (
	"net/http"

	"github.com/MonkyMars/gecho/utils"
)

// Success creates a 200 OK response with optional configuration.
func Success(w http.ResponseWriter, opts ...utils.ResponseOption) *utils.Response {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusOK),
		utils.WithMessage("Success"),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewOK(w, allOpts...)
}

// Created creates a 201 Created response with optional configuration.
func Created(w http.ResponseWriter, opts ...utils.ResponseOption) *utils.Response {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusCreated),
		utils.WithMessage("Resource Created"),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewOK(w, allOpts...)
}

// Accepted creates a 202 Accepted response with optional configuration.
func Accepted(w http.ResponseWriter, opts ...utils.ResponseOption) *utils.Response {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusAccepted),
		utils.WithMessage("Accepted"),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewOK(w, allOpts...)
}

// NoContent creates a 204 No Content response with optional configuration.
func NoContent(w http.ResponseWriter, opts ...utils.ResponseOption) *utils.Response {
	allOpts := []utils.ResponseOption{
		utils.WithStatus(http.StatusNoContent),
		utils.WithMessage("No Content"),
	}
	allOpts = append(allOpts, opts...)
	return utils.NewOK(w, allOpts...)
}
