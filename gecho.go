package gecho

import (
	"github.com/MonkyMars/gecho/clienterrors"
	"github.com/MonkyMars/gecho/pkg"
	"github.com/MonkyMars/gecho/servererrors"
)

// Exported pkg Functions
var Err = pkg.Err
var OK = pkg.OK

// Exported Client Error Functions
var BadRequest = clienterrors.BadRequest
var Unauthorized = clienterrors.Unauthorized
var Forbidden = clienterrors.Forbidden
var NotFound = clienterrors.NotFound

// Exported Server Error Functions
var InternalServerError = servererrors.InternalServerError
var ServiceUnavailable = servererrors.ServiceUnavailable

// Exported Types
type NewResponse = pkg.NewResponse
