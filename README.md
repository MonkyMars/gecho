# GECHO
### A simple response library for Go web frameworks using response writer.

<img src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fa-z-animals.com%2Fmedia%2F2022%2F12%2Fshutterstock_171751469.jpg&f=1&nofb=1&ipt=8b0fc9d6b628d7a8cb0055ab417dab5ecc1c3d3ec632bc5cb6c52207a5a5ec3e" alt="GECHO Logo" width="400" height="400">

## Overview
GECHO is a lightweight and easy-to-use response library designed to simplify the process of sending HTTP responses in Go web applications. It provides a fluent API for building JSON responses with automatic response handling.

## Installation
```bash
go get -u github.com/MonkyMars/gecho
```

## Response Structure
All responses follow this JSON structure:
```json
{
  "status": 200,
  "success": true,
  "message": "Success",
  "data": {...},
  "timestamp": "2024-01-01T12:00:00Z"
}
```

## Fluent API

### Building Responses
- `NewOK(w)` - Creates a success response builder (status 200)
- `NewErr(w)` - Creates an error response builder (status 500)

Chain methods for customization:
- `.WithMessage(string)` - Set response message
- `.WithData(any)` - Set response data (ignored for errors)
- `.WithStatus(int)` - Set HTTP status code
- `.Send()` - Send the response (required)

```go
// Fluent API example
gecho.NewOK(w).WithMessage("User created").WithData(user).WithStatus(http.StatusOK).Send()
```

## Success Functions

### Success(w, data)
Returns 200 OK with data and "Success" message.
```go
user := map[string]string{"name": "John"}
gecho.Success(w).WithData(user).Send()
```

### Created(w, data)
Returns 201 Created with data and "Resource Created" message.
```go
gecho.Created(w).WithData(newUser).Send()
```

## Client Error Functions

### BadRequest(w)
Returns 400 Bad Request with "Bad Request" message.
```go
gecho.BadRequest(w).Send()
// Or customize: gecho.BadRequest(w).WithMessage("Invalid input").Send()
```

### Unauthorized(w)
Returns 401 Unauthorized with "Unauthorized" message.
```go
gecho.Unauthorized(w).Send()
```

### Forbidden(w)
Returns 403 Forbidden with "Forbidden" message.
```go
gecho.Forbidden(w).Send()
```

### NotFound(w)
Returns 404 Not Found with "Not Found" message.
```go
gecho.NotFound(w).Send()
```

### MethodNotAllowed(w)
Returns 405 Method Not Allowed with "Method Not Allowed" message.
```go
gecho.MethodNotAllowed(w).Send()
```

## Server Error Functions

### InternalServerError(w)
Returns 500 Internal Server Error with "Internal Server Error" message.
```go
gecho.InternalServerError(w).Send()
```

### ServiceUnavailable(w)
Returns 503 Service Unavailable with "Service Unavailable" message.
```go
gecho.ServiceUnavailable(w).Send()
```

## Built-in Handlers

### Method Validation
Use `Handlers.HandleMethod()` to validate HTTP methods:
```go
func handler(w http.ResponseWriter, r *http.Request) {
    if err := gecho.Handlers.HandleMethod(w, r, "POST"); err != nil {
        err.Send() // Automatically sends 405 Method Not Allowed
        return
    }
    // Handle POST request
}
```

## Complete Example
```go
package main

import (
    "net/http"
    "github.com/MonkyMars/gecho"
)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
    // Method validation
    if err := gecho.Handlers.HandleMethod(w, r, "GET"); err != nil {
        err.Send()
        return
    }

    user := map[string]interface{}{
        "id":   1,
        "name": "John Doe",
        "email": "john@example.com",
    }

    gecho.Success(w).WithData(user).Send()
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
    if err := gecho.Handlers.HandleMethod(w, r, "POST"); err != nil {
        err.Send()
        return
    }

    // Validation logic here...
    if validationFailed {
        gecho.BadRequest(w).WithMessage("Invalid user data").Send()
        return
    }

    newUser := map[string]interface{}{"id": 2, "name": "Jane Doe"}
    gecho.Created(w).WithData(newUser).Send()
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
    gecho.InternalServerError(w).WithMessage("Something went wrong").Send()
}

func main() {
    http.HandleFunc("/user", getUserHandler)
    http.HandleFunc("/user/create", createUserHandler)
    http.HandleFunc("/error", errorHandler)
    http.ListenAndServe(":8080", nil)
}
```

## Features
- Fluent API for building responses
- Automatic JSON marshaling
- Consistent response structure
- Built-in HTTP status code handlers
- Method validation helpers
- Manual response control with `.Send()`
- Thread-safe response building

## Missing Features?
If you find any features missing or have suggestions for improvements, please feel free to open an issue or submit a pull request!
