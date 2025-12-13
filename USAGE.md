# GEcho Usage Guide

GEcho is a lightweight HTTP response handling library for Go that uses a functional options pattern similar to structured logging.

## Installation

```go
go get github.com/MonkyMars/gecho
```

## Quick Start

### Basic Error Response

```go
import (
    "net/http"
    "github.com/MonkyMars/gecho"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Send a 404 Not Found error
    gecho.NotFound(w, gecho.Send())
}
```

### Basic Success Response

```go
import (
    "github.com/MonkyMars/gecho/success"
    "github.com/MonkyMars/gecho/utils"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Send a 200 OK response
    gecho.Success(w, gecho.Send())
}
```

## Functional Options Pattern

GEcho uses a functional options pattern that allows you to chain configuration options:

### Available Options

- `utils.WithData(data any)` - Set response data
- `utils.WithMessage(message string)` - Override the default message
- `utils.WithStatus(status int)` - Override the default status code
- `utils.Send()` - Send the response immediately

### Response with Data

```go
func getUserHandler(w http.ResponseWriter, r *http.Request) {
    user := User{
        ID:       123,
        Username: "johndoe",
        Email:    "john@example.com",
    }
    
    gecho.Success(w, 
        gecho.WithData(user),
        gecho.Send(),
    )
}
```

### Custom Message

```go
func deleteHandler(w http.ResponseWriter, r *http.Request) {
    gecho.NoContent(w,
        gecho.WithMessage("Resource deleted successfully"),
        gecho.Send(),
    )
}
```

### Error with Validation Details

```go
func createUserHandler(w http.ResponseWriter, r *http.Request) {
    validationErrors := map[string]string{
        "email":    "Invalid email format",
        "password": "Password must be at least 8 characters",
    }
    
    gecho.BadRequest(w,
        gecho.WithData(validationErrors),
        gecho.Send(),
    )
}
```

### Multiple Options Combined

```go
func customResponseHandler(w http.ResponseWriter, r *http.Request) {
    gecho.Unauthorized(w,
        gecho.WithMessage("Your session has expired"),
        gecho.WithStatus(http.StatusUnauthorized),
        gecho.Send(),
    )
}
```

## Error Responses

All error functions are in the `errors` package:

### Client Errors (4xx)

```go
// 400 Bad Request
gecho.BadRequest(w, gecho.WithData(validationErrors), gecho.Send())

// 401 Unauthorized
gecho.Unauthorized(w, gecho.Send())

// 403 Forbidden
gecho.Forbidden(w, gecho.WithMessage("Access denied to this resource"), gecho.Send())

// 404 Not Found
gecho.NotFound(w, gecho.Send())

// 405 Method Not Allowed
gecho.MethodNotAllowed(w, gecho.Send())

// 409 Conflict
gecho.Conflict(w, gecho.WithMessage("Resource already exists"), gecho.Send())
```

### Server Errors (5xx)

```go
// 500 Internal Server Error
gecho.InternalServerError(w, gecho.Send())

// 503 Service Unavailable
gecho.ServiceUnavailable(w, gecho.WithMessage("Maintenance mode"), gecho.Send())
```

## Success Responses

All success functions are in the `success` package:

```go
// 200 OK
gecho.Success(w, gecho.WithData(data), gecho.Send())

// 201 Created
gecho.Created(w, gecho.WithData(newResource), gecho.Send())

// 202 Accepted
gecho.Accepted(w, gecho.WithMessage("Request queued for processing"), gecho.Send())

// 204 No Content
success.NoContent(w, gecho.Send())
```

### Method Handler

```go
import (
    "github.com/MonkyMars/gecho/handlers"
)

func protectedHandler(w http.ResponseWriter, r *http.Request) {
    h := handlers.NewHandlers()
    
    // Only allow POST requests
    if err := h.HandleMethod(w, r, http.MethodPost); err != nil {
        return // Error already sent
    }
    
    // Process POST request
    gecho.Success(w,
        gecho.WithMessage("Request processed"),
        gecho.Send(),
    )
}
```

## Response Format

All responses follow a consistent JSON structure:

```json
{
  "status": 200,
  "success": true,
  "message": "Success",
  "data": {
    "id": 123,
    "username": "johndoe"
  },
  "timestamp": "2024-01-15T10:30:45.123Z"
}
```

### Error Response

```json
{
  "status": 400,
  "success": false,
  "message": "Bad request",
  "data": <...>,
  "timestamp": "2024-01-15T10:30:45.123Z"
}
```

## Without Send()

If you don't include `gecho.Send()`, the response won't be sent automatically. This is useful when you want to configure the response but send it later or conditionally:

```go
func conditionalSend(w http.ResponseWriter, r *http.Request) {
    // This won't send anything
    err := gecho.Success(w,
        gecho.WithData(map[string]string{"key": "value"}),
        // Note: no Send() call
    )
    
    // err will be nil, but nothing is written to the response
}
```

## Direct Response Functions

You can also use the lower-level functions directly:

```go
// Success response
gecho.NewOK(w,
    gecho.WithMessage("Custom success"),
    gecho.WithData(myData),
    gecho.WithStatus(http.StatusOK),
    gecho.Send(),
)

// Error response
gecho.NewErr(w,
    gecho.WithMessage("Custom error"),
    gecho.WithStatus(http.StatusBadRequest),
    gecho.Send(),
)
```
