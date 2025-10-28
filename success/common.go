package success

import (
	"net/http"

	"github.com/MonkyMars/gecho/utils"
)

// Success returns a ResponseBuilder for 200 OK responses
// Use Send() to send the response with the default values
func Success(w http.ResponseWriter) *utils.ResponseBuilder {
	return utils.NewOK(w).WithMessage("Success")
}

// Created returns a ResponseBuilder for 201 Created responses
// Use Send() to send the response with the default values
func Created(w http.ResponseWriter) *utils.ResponseBuilder {
	return utils.NewOK(w).WithStatus(http.StatusCreated).WithMessage("Resource Created")
}

// Accepted returns a ResponseBuilder for 202 Accepted responses
// Use Send() to send the response with the default values
func Accepted(w http.ResponseWriter) *utils.ResponseBuilder {
	return utils.NewOK(w).WithStatus(http.StatusAccepted).WithMessage("Accepted")
}

// NoContent returns a ResponseBuilder for 204 No Content responses
// Use Send() to send the response with the default values
func NoContent(w http.ResponseWriter) *utils.ResponseBuilder {
	return utils.NewOK(w).WithStatus(http.StatusNoContent).WithMessage("No Content")
}
