package utils

import (
	"log"
	"os"
)

type logger struct{}

func Logger() *logger {
	return &logger{}
}

var (
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	initialized   bool
)

var logLevels = map[string]uint{
	"debug": 1,
	"info":  2,
	"warn":  3,
	"error": 4,
}

var currentLevel uint

// Init initializes the logger with configuration
// logLevel can be "info", "warn", or "error"
// "info" meaning all logs are shown
// "warn" meaning only warnings and errors are shown
// "error" meaning only errors are shown
// If an invalid log level is provided, the function will panic
func (l *logger) InitLogger(logLevel string) *logger {
	if initialized {
		return l
	}

	// Create default loggers that will be reconfigured later
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ltime|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "WARN: ", log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERR: ", log.Ltime|log.Lshortfile)

	if logLevel == "" {
		logLevel = "info"
	}

	if logLevel != "info" && logLevel != "warn" && logLevel != "error" {
		panic("invalid log level: " + logLevel)
	}

	currentLevel = logLevels[logLevel]
	if currentLevel == 0 {
		currentLevel = logLevels["info"]
	}

	initialized = true

	return l
}

func (l *logger) Info(v ...any) {
	if !initialized {
		panic("logger not initialized")
	}
	if currentLevel <= logLevels["info"] {
		infoLogger.Println(v...)
	}
}

func (l *logger) Warn(v ...any) {
	if !initialized {
		panic("logger not initialized")
	}
	if currentLevel <= logLevels["warn"] {
		warningLogger.Println(v...)
	}
}

func (l *logger) Err(v ...any) {
	if !initialized {
		panic("logger not initialized")
	}
	if currentLevel <= logLevels["error"] {
		errorLogger.Println(v...)
	}
}
