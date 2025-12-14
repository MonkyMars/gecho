package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MonkyMars/gecho"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// In-memory user storage for demo
var users = map[int]User{
	1: {ID: 1, Username: "alice", Email: "alice@example.com"},
	2: {ID: 2, Username: "bob", Email: "bob@example.com"},
}
var nextID = 3

var logger *gecho.Logger

func init() {
	logger = gecho.NewDefaultLogger()
}

func main() {
	// Create a new ServeMux for routing
	mux := http.NewServeMux()
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users/", userByIDHandler)
	mux.HandleFunc("/health", healthHandler)

	// Wrap the mux with logging middleware
	logger := gecho.NewLogger(gecho.NewConfig(gecho.WithShowCaller(false)))
	loggedHandler := gecho.Handlers.HandleLogging(mux, logger)

	fmt.Println("Server starting on :8080")
	fmt.Println("Try these endpoints:")
	fmt.Println("  GET  http://localhost:8080/users")
	fmt.Println("  GET  http://localhost:8080/users/1")
	fmt.Println("  POST http://localhost:8080/users")
	fmt.Println("  GET  http://localhost:8080/health")

	log.Fatal(http.ListenAndServe(":8080", loggedHandler))
}

// Health check endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET
	if err := gecho.Handlers.HandleMethod(w, r, http.MethodGet); err != nil {
		return
	}

	gecho.Success(w,
		gecho.WithData(map[string]string{
			"status":  "healthy",
			"version": "1.0.0",
		}),
		gecho.Send(),
	)
}

// List all users or create a new user
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listUsers(w)
	case http.MethodPost:
		createUser(w, r)
	default:
		gecho.MethodNotAllowed(w,
			gecho.WithMessage(fmt.Sprintf("Method %s not allowed", r.Method)),
			gecho.Send(),
		)
	}
}

// List all users
func listUsers(w http.ResponseWriter) {
	userList := make([]User, 0, len(users))
	for _, user := range users {
		userList = append(userList, user)
	}

	logger.Info("Listing all users", gecho.Field("count", len(userList)))

	gecho.Success(w,
		gecho.WithData(map[string]any{
			"users": userList,
			"count": len(userList),
		}),
		gecho.Send(),
	)
}

// Create a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		gecho.BadRequest(w,
			gecho.WithMessage("Invalid request body"),
			gecho.Send(),
		)
		return
	}

	// Validation
	validationErrors := make(map[string]string)
	if req.Username == "" {
		validationErrors["username"] = "Username is required"
	}
	if req.Email == "" {
		validationErrors["email"] = "Email is required"
	}

	if len(validationErrors) > 0 {
		gecho.BadRequest(w,
			gecho.WithMessage("Validation failed"),
			gecho.WithData(validationErrors),
			gecho.Send(),
		)
		return
	}

	// Check if user already exists
	for _, user := range users {
		if user.Email == req.Email {
			gecho.Conflict(w,
				gecho.WithMessage("User with this email already exists"),
				gecho.Send(),
			)
			return
		}
	}

	// Create new user
	newUser := User{
		ID:       nextID,
		Username: req.Username,
		Email:    req.Email,
	}
	users[nextID] = newUser
	nextID++

	gecho.Created(w,
		gecho.WithData(newUser),
		gecho.WithMessage("User created successfully"),
		gecho.Send(),
	)
}

// Get user by ID
func userByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET
	if err := gecho.Handlers.HandleMethod(w, r, http.MethodGet); err != nil {
		return
	}

	// Extract ID from path (simple parsing for demo)
	var id int
	_, err := fmt.Sscanf(r.URL.Path, "/users/%d", &id)
	if err != nil {
		gecho.BadRequest(w,
			gecho.WithMessage("Invalid user ID"),
			gecho.Send(),
		)
		return
	}

	// Find user
	user, exists := users[id]
	if !exists {
		gecho.NotFound(w,
			gecho.WithMessage(fmt.Sprintf("User with ID %d not found", id)),
			gecho.Send(),
		)
		return
	}

	gecho.Success(w,
		gecho.WithData(user),
		gecho.Send(),
	)
}
