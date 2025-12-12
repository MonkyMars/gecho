package utils

import (
	"log"
	"os"
)

type Logger struct{}

func GetLogger() *Logger {
	return &Logger{}
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
func InitLogger(logLevel ...string) *Logger {
	if initialized {
		return &Logger{}
	}

	// Create default loggers that will be reconfigured later
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ltime|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "WARN: ", log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERR: ", log.Ltime|log.Lshortfile)

	if len(logLevel) == 0 {
		logLevel = []string{"info"}
	}

	if logLevel[0] != "info" && logLevel[0] != "warn" && logLevel[0] != "error" {
		panic("invalid log level: " + logLevel[0])
	}

	currentLevel = logLevels[logLevel[0]]
	if currentLevel == 0 {
		currentLevel = logLevels["info"]
	}

	initialized = true

	return &Logger{}
}

func (l *Logger) Info(v ...any) {
	if !initialized {
		panic("logger not initialized")
	}
	if currentLevel <= logLevels["info"] {
		infoLogger.Println(v...)
	}
}

func (l *Logger) Warn(v ...any) {
	if !initialized {
		panic("logger not initialized")
	}
	if currentLevel <= logLevels["warn"] {
		warningLogger.Println(v...)
	}
}

func (l *Logger) Err(v ...any) {
	if !initialized {
		panic("logger not initialized")
	}
	if currentLevel <= logLevels["error"] {
		errorLogger.Println(v...)
	}
}
