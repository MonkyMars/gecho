package gecho

import (
	"github.com/MonkyMars/gecho/clienterrors"
	"github.com/MonkyMars/gecho/pkg"
	"github.com/MonkyMars/gecho/servererrors"
)

var Err = pkg.Err
var OK = pkg.OK

var BadRequest = clienterrors.BadRequest
var Unauthorized = clienterrors.Unauthorized
var Forbidden = clienterrors.Forbidden
var NotFound = clienterrors.NotFound

var InternalServerError = servererrors.InternalServerError
var ServiceUnavailable = servererrors.ServiceUnavailable
