package gecho

import (
	"github.com/MonkyMars/gecho/errors"
	"github.com/MonkyMars/gecho/success"
	"github.com/MonkyMars/gecho/utils"
)

// Exported fluent API Functions
var NewErr = utils.NewErr
var NewOK = utils.NewOK

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

// Exported built-in handlers
var Handlers = errors.NewHandlers()
