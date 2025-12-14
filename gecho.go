package gecho

import (
	"github.com/MonkyMars/gecho/errors"
	"github.com/MonkyMars/gecho/handlers"
	"github.com/MonkyMars/gecho/success"
	"github.com/MonkyMars/gecho/utils"
)

// Response option type
type ResponseOption = utils.ResponseOption

// Response option functions
var WithData = utils.WithData
var WithMessage = utils.WithMessage
var WithStatus = utils.WithStatus
var Send = utils.Send

// Exported fluent API Functions
var NewErr = utils.NewErr
var NewOK = utils.NewOK

// Exported Client Error Functions
var BadRequest = errors.BadRequest
var Unauthorized = errors.Unauthorized
var Forbidden = errors.Forbidden
var NotFound = errors.NotFound
var MethodNotAllowed = errors.MethodNotAllowed
var Conflict = errors.Conflict

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
var Field = utils.Field
var WithCallerSkip = utils.WithCallerSkip

// Logger config functions
var NewConfig = utils.NewConfig
var WithLogLevel = utils.WithLogLevel
var WithLogFormat = utils.WithLogFormat
var WithColorize = utils.WithColorize
var WithShowCaller = utils.WithShowCaller
var WithTimeFormat = utils.WithTimeFormat
var WithOutput = utils.WithOutput
var WithErrorOutput = utils.WithErrorOutput
var WithDefaultCallerSkip = utils.WithDefaultCallerSkip

// Logger types
type Logger = utils.Logger
type LoggerConfig = utils.Config
type LoggerOptions = utils.LoggerOptions

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
	LogFormatText   = utils.FormatText
	LogFormatJSON   = utils.FormatJSON
	LogFormatPretty = utils.FormatPretty
)
