# Gecho

![gecho banner](https://static.wixstatic.com/media/a002ff_126a2882368c4a9c9194eb59df5d2cc4~mv2.jpg/v1/fill/w_640,h_332,al_c,q_80,usm_0.66_1.00_0.01,enc_avif,quality_auto/a002ff_126a2882368c4a9c9194eb59df5d2cc4~mv2.jpg)


A lightweight JSON response builder for Go HTTP handlers.

## What it is

`gecho` provides a small, fluent API for building and sending consistent JSON HTTP responses from Go handlers (success and error helpers, customizable message/data/status, and built-in method validation helpers).

## Install

Requires Go toolchain.

```bash
go get github.com/MonkyMars/gecho
```

## Quick example

```go
// in an http.Handler:
gecho.Success(w).WithData(map[string]any{"id":1, "name":"Alice"}).WithMessage("OK").Send()

// convenience error
gecho.BadRequest(w).WithMessage("invalid input").Send()
```

## Project layout

- `gecho.go` — package entry points
- `errors/` — error response helpers
- `success/` — success response helpers
- `handlers/` — small built-in handlers (method validation)
- `utils/` — core builder and JSON logic

## Contributing

Contributions welcome:) Please open issues or pull requests and include a short description and tests for behavior changes.
