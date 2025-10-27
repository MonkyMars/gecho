package gecho

import (
	"github.com/MonkyMars/gecho/errors"
	"github.com/MonkyMars/gecho/pkg"
	"github.com/MonkyMars/gecho/success"
)

// Exported fluent API Functions
var NewErr = pkg.NewErr
var NewOK = pkg.NewOK

// Exported Client Error Functions
var BadRequest = errors.BadRequest
var Unauthorized = errors.Unauthorized
var Forbidden = errors.Forbidden
var NotFound = errors.NotFound
var MethodNotAllowed = errors.MethodNotAllowed

// Exported Server Error Functions
var InternalServerError = errors.InternalServerError
var ServiceUnavailable = errors.ServiceUnavailable

// Exported Success Functions
var Success = success.Success[any]
var Created = success.Created[any]
