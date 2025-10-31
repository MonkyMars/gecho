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

````markdown
# GECHO

A small, focused response-helper library for Go web handlers.

This repository provides a compact API to build and send JSON HTTP responses with a fluent builder. The implementation keeps response fields private while exposing a JSON representation via custom marshal/unmarshal logic.

Key packages:

- `github.com/MonkyMars/gecho` — root package that re-exports the public API (convenience surface).
- `github.com/MonkyMars/gecho/utils` — core ResponseBuilder and `NewResponse` type.
- `github.com/MonkyMars/gecho/errors` — common error response helpers (400, 401, 403, 404, 405, 500, 503).
- `github.com/MonkyMars/gecho/success` — common success helpers (200, 201, 202, 204).
- `github.com/MonkyMars/gecho/handlers` — small built-in handlers (method validation).

Installation

```bash
go get -u github.com/MonkyMars/gecho
```
````

Response JSON format

```json
{
  "status": 200,
  "success": true,
  "message": "Success",
  "data": null,
  "timestamp": "2025-10-31T12:00:00Z"
}
```

Notes

- The `NewResponse` struct fields are intentionally unexported. The package implements `MarshalJSON`/`UnmarshalJSON` so JSON encoding/decoding still produces the expected keys.
- When building error responses, any `.WithData(...)` value is ignored and the response `data` will be `null` in the JSON output.

Quick API

- gecho.NewOK(w) / gecho.NewErr(w) — start a response builder (success / error)
- gecho.Success(w), gecho.Created(w), gecho.Accepted(w), gecho.NoContent(w) — convenience success builders (re-exported from `success` package)
- gecho.BadRequest(w), gecho.Unauthorized(w), gecho.Forbidden(w), gecho.NotFound(w), gecho.MethodNotAllowed(w) — client errors
- gecho.InternalServerError(w), gecho.ServiceUnavailable(w) — server errors
- gecho.Handlers — built-in handler helpers (method validation)

Fluent builder methods (chainable)

- `.WithMessage(string)` — set response message
- `.WithData(any)` — set response data (ignored for error responses)
- `.WithStatus(int)` — set non-default status code
- `.Send()` — encode and write the response to the provided `http.ResponseWriter`

Example (method validation + success)

```go
package main

import (
    "net/http"
    "github.com/MonkyMars/gecho"
)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
    // Ensure only GET is allowed. HandleMethod returns a *ResponseBuilder when the
    // request method is not allowed (so you can customize or send it directly).
    if rb := gecho.Handlers.HandleMethod(w, r, http.MethodGet); rb != nil {
        _ = rb.Send() // sends a 405 Method Not Allowed by default
        return
    }

    user := map[string]any{"id": 1, "name": "John Doe"}
    _ = gecho.Success(w).WithData(user).Send()
}

func main() {
    http.HandleFunc("/user", getUserHandler)
    _ = http.ListenAndServe(":8080", nil)
}
```

Testing

```bash
# run all tests (verbose)
go test ./... -v

# run tests with coverage
go test ./... -cover
```

Contributing

- Open an issue or a pull request if you find bugs or want to add features.

License

- (Add license information here if you have one.)

```
- Consistent response structure
```
