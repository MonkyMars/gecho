package gecho

import (
	"github.com/MonkyMars/gecho/errors"
	"github.com/MonkyMars/gecho/handlers"
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
var Success = success.Success
var Created = success.Created

// Exported built-in handlers
var Handlers = handlers.NewHandlers()

// Logger exports
var NewLogger = utils.NewLogger
var NewDefaultLogger = utils.NewDefaultLogger
var DefaultLoggerConfig = utils.DefaultConfig
var ParseLogLevel = utils.ParseLevel

// Log levels
var (
	LogLevelDebug = utils.LevelDebug
	LogLevelInfo  = utils.LevelInfo
	LogLevelWarn  = utils.LevelWarn
	LogLevelError = utils.LevelError
	LogLevelFatal = utils.LevelFatal
)

// Log formats
var (
	LogFormatText = utils.FormatText
	LogFormatJSON = utils.FormatJSON
)
