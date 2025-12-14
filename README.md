# Gecho

A Go library for structured logging and consistent JSON HTTP responses.

## Install

```bash
go get github.com/MonkyMars/gecho
```

## What It Does

Gecho provides two main features:

1. **Structured Logger** - A thread-safe logger with multiple output formats and log levels
2. **HTTP Response Builder** - Consistent JSON responses for HTTP handlers

## Response Handling

### Basic Usage

```go
import "github.com/MonkyMars/gecho"

func handler(w http.ResponseWriter, r *http.Request) {
    // Success response
    gecho.Success(w, gecho.Send())
    
    // Success with data
    gecho.Success(w,
        gecho.WithData(map[string]any{"id": 1, "name": "Alice"}),
        gecho.Send(),
    )
    
    // Error response
    gecho.NotFound(w,
        gecho.WithMessage("User not found"),
        gecho.Send(),
    )
}
```

### Available Functions

**Success Responses:**
- `Success(w, opts...)` - 200 OK
- `Created(w, opts...)` - 201 Created
- `Accepted(w, opts...)` - 202 Accepted
- `NoContent(w, opts...)` - 204 No Content

**Error Responses:**
- `BadRequest(w, opts...)` - 400 Bad Request
- `Unauthorized(w, opts...)` - 401 Unauthorized
- `Forbidden(w, opts...)` - 403 Forbidden
- `NotFound(w, opts...)` - 404 Not Found
- `MethodNotAllowed(w, opts...)` - 405 Method Not Allowed
- `Conflict(w, opts...)` - 409 Conflict
- `InternalServerError(w, opts...)` - 500 Internal Server Error
- `ServiceUnavailable(w, opts...)` - 503 Service Unavailable

### Options

- `WithData(data any)` - Add data to response
- `WithMessage(msg string)` - Override default message
- `WithStatus(code int)` - Override default status code
- `Send()` - Send the response (required)

### Response Format

All responses return this JSON structure:

```json
{
  "status": 200,
  "success": true,
  "message": "Success",
  "data": {"id": 1, "name": "Alice"},
  "timestamp": "2024-01-15T10:30:45.123Z"
}
```

## Logger

### Basic Usage

```go
import "github.com/MonkyMars/gecho"

func main() {
    // Create logger with defaults
    logger := gecho.NewDefaultLogger()
    
    // Log messages
    logger.Info("Server starting")
    logger.Error("Failed to connect")
    
    // Log with fields
    logger.Info("User logged in",
        gecho.Field("user_id", 123),
        gecho.Field("ip", "192.168.1.1"),
    )
}
```

### Configuration

```go
// Custom configuration
config := gecho.NewConfig(
    gecho.WithLogLevel(gecho.LevelDebug),
    gecho.WithLogFormat(gecho.FormatJSON),
    gecho.WithShowCaller(true),
)
logger := gecho.NewLogger(config)

// Change level at runtime
logger.SetLevel(gecho.ParseLogLevel("debug"))
```

### Configuration Options

- `WithLogLevel(level Level)` - Set minimum log level (default: `LevelInfo`)
- `WithLogFormat(format Format)` - Set output format (default: `FormatPretty`)
- `WithColorize(bool)` - Enable/disable colored output (default: auto-detected)
- `WithShowCaller(bool)` - Show/hide file and line number (default: `true`)
- `WithTimeFormat(string)` - Custom time format (default: `"2006-01-02 15:04:05.000"`)
- `WithOutput(io.Writer)` - Set output destination (default: `os.Stdout`)
- `WithErrorOutput(io.Writer)` - Set error output destination (default: `os.Stderr`)
- `WithDefaultCallerSkip(int)` - Adjust call stack depth for caller info (default: `2`)

### Log Levels

- `LevelDebug` - Debug messages
- `LevelInfo` - Informational messages
- `LevelWarn` - Warning messages
- `LevelError` - Error messages
- `LevelFatal` - Fatal errors (exits program)

### Output Formats

- `FormatText` - Plain text with fields
- `FormatJSON` - JSON output
- `FormatPretty` - Colored output with parentheses format (default)

### Persistent Fields

```go
// Create logger with persistent fields
requestLogger := logger.WithFields(map[string]any{
    "request_id": "abc123",
    "user_id": 456,
})

requestLogger.Info("Processing request") // All logs include request_id and user_id
```

### HTTP Logging Middleware

```go
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handler)
    
    logger := gecho.NewDefaultLogger()
    loggedHandler := gecho.Handlers.HandleLogging(mux, logger)
    
    http.ListenAndServe(":8080", loggedHandler)
}
```

Logs include method, path, status, duration, and remote address.

## Method Validation

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Only allow POST requests
    if err := gecho.Handlers.HandleMethod(w, r, http.MethodPost); err != nil {
        return // Error response already sent
    }
    
    // Handle POST request
}
```

## Full Example

```go
package main

import (
    "net/http"
    "github.com/MonkyMars/gecho"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/users", getUsers)
    
    logger := gecho.NewDefaultLogger()
    loggedHandler := gecho.Handlers.HandleLogging(mux, logger)
    
    http.ListenAndServe(":8080", loggedHandler)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    if err := gecho.Handlers.HandleMethod(w, r, http.MethodGet); err != nil {
        return
    }
    
    users := []map[string]any{
        {"id": 1, "name": "Alice"},
        {"id": 2, "name": "Bob"},
    }
    
    gecho.Success(w,
        gecho.WithData(map[string]any{"users": users}),
        gecho.Send(),
    )
}
```

## Project Structure

- `gecho.go` - Main package exports
- `errors/` - Error response functions
- `success/` - Success response functions
- `handlers/` - HTTP middleware and utilities
- `utils/` - Core response builder and logger

## Contributing

Contributions are welcome. Open an issue or pull request with a description and tests for changes.

## Note

This library is provided as-is. Breaking changes may occur between versions.