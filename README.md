# Gecho

![gecho banner](https://static.wixstatic.com/media/a002ff_126a2882368c4a9c9194eb59df5d2cc4~mv2.jpg/v1/fill/w_640,h_332,al_c,q_80,usm_0.66_1.00_0.01,enc_avif,quality_auto/a002ff_126a2882368c4a9c9194eb59df5d2cc4~mv2.jpg)


A lightweight JSON response builder for Go HTTP handlers.

## What it is

`gecho` provides a lightweight, functional options-based API for building and sending consistent JSON HTTP responses from Go handlers. It offers success and error helpers, customizable message/data/status, and built-in method validation helpers using a pattern similar to structured logging.

## Install

Requires Go toolchain.

```bash
go get github.com/MonkyMars/gecho
```

## Quick example

```go
import (
    "github.com/MonkyMars/gecho"
)

// Success response with data
gecho.Success(w, 
    gecho.WithData(map[string]any{"id": 1, "name": "Alice"}),
    gecho.Send(),
)

// Error response with custom message
gecho.BadRequest(w, 
    gecho.WithMessage("invalid input"),
    gecho.Send(),
)

// Multiple options
gecho.NotFound(w,
    gecho.WithMessage("User not found"),
    gecho.WithData(map[string]string{"resource": "users", "id": "123"}),
    gecho.Send(),
)
```

## Response Format

All responses return a consistent JSON structure:

```json
{
  "status": 200,
  "success": true,
  "message": "Success",
  "data": {
    "id": 1,
    "name": "Alice"
  },
  "timestamp": "2024-01-15T10:30:45.123Z"
}
```

## Available Options

- `utils.WithData(data any)` - Set response data
- `utils.WithMessage(message string)` - Override the default message
- `utils.WithStatus(status int)` - Override the default status code
- `utils.Send()` - Send the response immediately

## Common Usage Patterns

### Success Responses

```go
// 200 OK
success.Success(w, utils.WithData(userData), utils.Send())

// 201 Created
success.Created(w, utils.WithData(newResource), utils.Send())

// 202 Accepted
success.Accepted(w, utils.Send())

// 204 No Content
success.NoContent(w, utils.Send())
```

### Error Responses

```go
// 400 Bad Request
errors.BadRequest(w, utils.WithData(validationErrors), utils.Send())

// 401 Unauthorized
errors.Unauthorized(w, utils.Send())

// 403 Forbidden
errors.Forbidden(w, utils.Send())

// 404 Not Found
errors.NotFound(w, utils.Send())

// 500 Internal Server Error
errors.InternalServerError(w, utils.Send())
```

### Full Example

```go
func getUserHandler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    
    if id == "" {
        gecho.BadRequest(w,
            gecho.WithMessage("Missing user ID"),
            gecho.Send(),
        )
        return
    }
    
    user, err := findUser(id)
    if err != nil {
        gecho.NotFound(w, utils.Send())
        return
    }
    
    gecho.Success(w,
        gecho.WithData(user),
        gecho.Send(),
    )
}
```

For more detailed examples and documentation, see [USAGE.md](USAGE.md).

## Project layout

- `gecho.go` — package entry points
- `errors/` — error response helpers
- `success/` — success response helpers
- `handlers/` — small built-in handlers (method validation)
- `utils/` — core builder and JSON logic

## Contributing

Contributions welcome:) Please open issues or pull requests and include a short description and tests for behavior changes.
