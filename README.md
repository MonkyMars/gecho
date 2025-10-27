# GECHO
### A simple response library for Go web frameworks using response writer.

<img src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fa-z-animals.com%2Fmedia%2F2022%2F12%2Fshutterstock_171751469.jpg&f=1&nofb=1&ipt=8b0fc9d6b628d7a8cb0055ab417dab5ecc1c3d3ec632bc5cb6c52207a5a5ec3e" alt="GECHO Logo" width="400" height="400">

## Overview
GECHO is a lightweight and easy-to-use response library designed to simplify the process of sending HTTP responses in Go web applications. It provides a set of helper functions to send JSON with ease.

## Installation
To install GECHO, use the following command:
```bash
go get -u github.com/MonkyMars/gecho
```

## Usage
Here's a simple example of how to use GECHO in your Go web application:
```go
package main

import (
	"net/http"
	"github.com/MonkyMars/gecho"
)

// Success
func successHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, GECHO!"}
	gecho.Success(w, response).Send()
}

// Error
func errorHandler(w http.ResponseWriter, r *http.Request) {
	gecho.Error(w, "An error occurred").Send()
}

func main() {
	http.HandleFunc("/success", handler)
	http.HandleFunc("/error", errorHandler)
	http.ListenAndServe(":8080", nil)
}
```


## Missing Features?
If you find any features missing or have suggestions for improvements, please feel free to open an issue or submit a pull request!
