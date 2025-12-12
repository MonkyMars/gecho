package utils_test

import (
	"bytes"
	"os"

	"github.com/MonkyMars/gecho/utils"
)

// Example_basicUsage demonstrates basic logger usage
func Example_basicUsage() {
	// Create a logger with default settings
	logger := utils.NewDefaultLogger()

	// Log messages at different levels
	logger.Info("Application started")
	logger.Warn("This is a warning")
	logger.Error("An error occurred")

	// Debug messages won't show (default level is Info)
	logger.Debug("This won't appear")
}

// Example_structuredLogging demonstrates structured key-value logging
func Example_structuredLogging() {
	logger := utils.NewDefaultLogger()

	// Log with structured key-value pairs
	logger.Info("User logged in",
		"user_id", 12345,
		"username", "john_doe",
		"ip_address", "192.168.1.1",
	)

	logger.Error("Database connection failed",
		"error", "connection timeout",
		"host", "db.example.com",
		"port", 5432,
		"retry_count", 3,
	)
}

// Example_customConfiguration demonstrates custom logger configuration
func Example_customConfiguration() {
	// Create custom configuration
	config := utils.Config{
		Level:       utils.LevelDebug, // Show all logs including debug
		Format:      utils.FormatText, // Human-readable format
		Output:      os.Stdout,        // Standard output
		ErrorOutput: os.Stderr,        // Errors to stderr
		Colorize:    true,             // Enable colors
		ShowCaller:  true,             // Show file:line
		TimeFormat:  "15:04:05.000",   // Custom time format
	}

	logger := utils.NewLogger(config)
	logger.Info("Custom logger initialized")
}

// Example_jsonFormat demonstrates JSON output format
func Example_jsonFormat() {
	config := utils.DefaultConfig()
	config.Format = utils.FormatJSON

	logger := utils.NewLogger(config)

	logger.Info("Request processed",
		"method", "GET",
		"path", "/api/users",
		"status", 200,
		"duration_ms", 45,
	)
	// Output: {"timestamp":"...","level":"INFO","message":"Request processed","caller":"...","fields":{"method":"GET","path":"/api/users","status":200,"duration_ms":45}}
}

// Example_contextualLogging demonstrates adding persistent fields
func Example_contextualLogging() {
	// Create base logger
	logger := utils.NewDefaultLogger()

	// Create request-scoped logger with request ID
	requestLogger := logger.WithField("request_id", "abc-123")

	// All logs from this logger include the request_id
	requestLogger.Info("Starting request processing")
	requestLogger.Info("Database query executed", "rows", 42)
	requestLogger.Info("Request completed", "status", 200)

	// Add multiple fields at once
	userLogger := requestLogger.WithFields(map[string]any{
		"user_id":  456,
		"username": "alice",
	})

	userLogger.Info("User action performed", "action", "update_profile")
}

// Example_logLevels demonstrates different log levels
func Example_logLevels() {
	logger := utils.NewDefaultLogger()

	// Set minimum log level
	logger.SetLevel(utils.LevelWarn)

	// These won't show (below warn level)
	logger.Debug("Debug message")
	logger.Info("Info message")

	// These will show
	logger.Warn("Warning message")
	logger.Error("Error message")
}

// Example_parseLevel demonstrates parsing log level from string
func Example_parseLevel() {
	config := utils.DefaultConfig()

	// Parse level from environment variable or config file
	levelStr := os.Getenv("LOG_LEVEL")
	if levelStr == "" {
		levelStr = "info"
	}
	config.Level = utils.ParseLevel(levelStr)

	logger := utils.NewLogger(config)
	logger.Info("Logger initialized", "level", levelStr)
}

// Example_customOutput demonstrates writing logs to custom destinations
func Example_customOutput() {
	// Write to a buffer instead of stdout
	var buf bytes.Buffer

	config := utils.DefaultConfig()
	config.Output = &buf
	config.Colorize = false // Disable colors for buffer output

	logger := utils.NewLogger(config)
	logger.Info("Test message", "key", "value")

	// The log is now in buf, which could be written to a file, sent to a service, etc.
}

// Example_httpMiddleware demonstrates using logger in HTTP middleware
func Example_httpMiddleware() {
	// Create logger for the application
	logger := utils.NewDefaultLogger()

	// In your HTTP handler/middleware
	handleRequest := func(requestID string, userID int) {
		// Create request-scoped logger
		reqLogger := logger.WithFields(map[string]any{
			"request_id": requestID,
			"user_id":    userID,
		})

		reqLogger.Info("Processing request")

		// Pass reqLogger to other functions
		// All logs will include request_id and user_id
		processBusinessLogic(reqLogger)

		reqLogger.Info("Request completed", "status", 200)
	}

	handleRequest("req-123", 456)
}

func processBusinessLogic(logger *utils.Logger) {
	logger.Debug("Executing business logic")
	logger.Info("Database query executed", "duration_ms", 23)
}

// Example_errorHandling demonstrates logging errors
func Example_errorHandling() {
	logger := utils.NewDefaultLogger()

	// Simulate an error
	err := performOperation()
	if err != nil {
		logger.Error("Operation failed",
			"error", err.Error(),
			"operation", "database_query",
			"retry_attempted", true,
		)
	}
}

func performOperation() error {
	return nil
}

// Example_productionSetup demonstrates a typical production configuration
func Example_productionSetup() {
	// Production logger configuration
	config := utils.Config{
		Level:       utils.ParseLevel(getEnv("LOG_LEVEL", "info")),
		Format:      parseFormat(getEnv("LOG_FORMAT", "json")),
		Output:      os.Stdout,
		ErrorOutput: os.Stderr,
		Colorize:    false, // Disable colors for log aggregation
		ShowCaller:  true,
		TimeFormat:  "2006-01-02T15:04:05.000Z07:00", // ISO 8601
	}

	logger := utils.NewLogger(config)

	// Add application metadata
	appLogger := logger.WithFields(map[string]any{
		"app":     "myapp",
		"version": "1.2.3",
		"env":     getEnv("ENVIRONMENT", "production"),
	})

	appLogger.Info("Application started")
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func parseFormat(s string) utils.Format {
	if s == "json" {
		return utils.FormatJSON
	}
	return utils.FormatText
}

// Example_multipleLoggers demonstrates using different loggers for different purposes
func Example_multipleLoggers() {
	// Access logger with INFO level
	accessConfig := utils.DefaultConfig()
	accessConfig.Level = utils.LevelInfo
	accessLogger := utils.NewLogger(accessConfig).WithField("logger", "access")

	// Error logger with ERROR level only
	errorConfig := utils.DefaultConfig()
	errorConfig.Level = utils.LevelError
	errorLogger := utils.NewLogger(errorConfig).WithField("logger", "error")

	// Debug logger for development
	debugConfig := utils.DefaultConfig()
	debugConfig.Level = utils.LevelDebug
	debugLogger := utils.NewLogger(debugConfig).WithField("logger", "debug")

	// Use different loggers for different purposes
	accessLogger.Info("User accessed page", "path", "/home")
	errorLogger.Error("Critical error occurred", "code", 500)
	debugLogger.Debug("Variable state", "x", 42, "y", 100)
}
