package success

import (
	"net/http"

	"github.com/MonkyMars/gecho/pkg"
)

// Success returns a ResponseBuilder for 200 OK responses
// Use Send() to send the response with the default values
func Success[T any](w http.ResponseWriter, data T) *pkg.ResponseBuilder {
	return pkg.NewOK(w).WithStatus(http.StatusOK).
		WithData(data).
		WithMessage("Success")
}

// Created returns a ResponseBuilder for 201 Created responses
// Use Send() to send the response with the default values
func Created[T any](w http.ResponseWriter, data T) *pkg.ResponseBuilder {
	return pkg.NewOK(w).WithStatus(http.StatusCreated).
		WithData(data).
		WithMessage("Resource Created")
}

// Accepted returns a ResponseBuilder for 202 Accepted responses
// Use Send() to send the response with the default values
func Accepted(w http.ResponseWriter) *pkg.ResponseBuilder {
	return pkg.NewOK(w).WithStatus(http.StatusAccepted).WithMessage("Accepted")
}

// NoContent returns a ResponseBuilder for 204 No Content responses
// Use Send() to send the response with the default values
func NoContent(w http.ResponseWriter) *pkg.ResponseBuilder {
	return pkg.NewOK(w).WithStatus(http.StatusNoContent).WithMessage("No Content")
}
