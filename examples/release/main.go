// Release example for the error-aware response constructors and HTTP middleware.
//
// Run from this directory with:
//
//	go run .
//
// Then try:
//
//	curl -i http://localhost:8080/health
//	curl -i http://localhost:8080/ready
//	curl -i -X POST http://localhost:8080/users
//	curl -i -X DELETE http://localhost:8080/users
package main

import (
	"log"
	"net/http"

	"github.com/MonkyMars/gecho"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health)
	mux.HandleFunc("/ready", ready)
	mux.HandleFunc("/users", createUser)

	logger := gecho.NewLogger(gecho.NewConfig(
		gecho.WithShowCaller(false),
	))
	handler := gecho.Handlers.HandleLogging(mux, logger)

	log.Println("release example listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		if err := gecho.MethodNotAllowed(w, gecho.WithMessage("only GET is supported")).Send(); err != nil {
			log.Printf("write method error: %v", err)
		}
		return
	}

	// Immediate sends return write and encoding errors.
	if err := gecho.Success(w,
		gecho.WithData(map[string]string{"status": "healthy"}),
		gecho.WithHeader("X-Release", "error-aware-responses"),
	).Send(); err != nil {
		log.Printf("write health response: %v", err)
	}
}

func ready(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		_ = gecho.MethodNotAllowed(w).Send()
		return
	}

	// NoContent now emits a protocol-compliant 204 response with no body.
	if err := gecho.NoContent(w).Send(); err != nil {
		log.Printf("write readiness response: %v", err)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		if err := gecho.MethodNotAllowed(w,
			gecho.WithStatus(http.StatusMethodNotAllowed),
			gecho.WithMessage("use POST to create a user"),
		).Send(); err != nil {
			log.Printf("write method error: %v", err)
		}
		return
	}

	if err := gecho.Created(w,
		gecho.WithData(map[string]string{"id": "user-123"}),
	).Send(); err != nil {
		log.Printf("write user response: %v", err)
	}
}
