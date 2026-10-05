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
var nextID = len(users) + 1

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
	logLevel := gecho.ParseLogLevel("debug")
	logger := gecho.NewLogger(gecho.NewConfig(gecho.WithShowCaller(true)))
	logger.SetLevel(logLevel)
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

	if err := gecho.Success(w,
		gecho.WithData(map[string]string{
			"status":  "healthy",
			"version": "1.0.0",
		}),
		gecho.WithMessage("Health check passed"),
	).Send(); err != nil {
		logger.Error("send health response", gecho.Field("error", err))
	}
}

// List all users or create a new user
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listUsers(w)
	case http.MethodPost:
		createUser(w, r)
	default:
		if err := gecho.MethodNotAllowed(w,
			gecho.WithMessage(fmt.Sprintf("Method %s not allowed", r.Method)),
		).Send(); err != nil {
			logger.Error("send method error", gecho.Field("error", err))
		}
	}
}

// List all users
func listUsers(w http.ResponseWriter) {
	userList := make([]User, 0, len(users))
	for _, user := range users {
		userList = append(userList, user)
	}

	logger.Info("Listing all users", gecho.Field("count", len(userList)))

	responseData := map[string]any{
		"users": userList,
		"count": len(userList),
	}
	if err := gecho.Success(w,
		gecho.WithData(responseData),
		gecho.WithMessage("Users retrieved successfully"),
	).Send(); err != nil {
		logger.Error("send response", gecho.Field("error", err))
	}
}

// Create a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		if err := gecho.BadRequest(w,
			gecho.WithMessage("Invalid request body"),
		).Send(); err != nil {
			logger.Error("send request error", gecho.Field("error", err))
		}
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
		if err := gecho.BadRequest(w,
			gecho.WithMessage("Validation failed"),
			gecho.WithData(validationErrors),
		).Send(); err != nil {
			logger.Error("send validation error", gecho.Field("error", err))
		}
		return
	}

	// Check if user already exists
	for _, user := range users {
		if user.Email == req.Email {
			if err := gecho.Conflict(w,
				gecho.WithMessage("User with this email already exists"),
			).Send(); err != nil {
				logger.Error("send conflict error", gecho.Field("error", err))
			}
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

	if err := gecho.Created(w,
		gecho.WithData(newUser),
		gecho.WithMessage("User created successfully"),
	).Send(); err != nil {
		logger.Error("send created response", gecho.Field("error", err))
	}
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
		if err := gecho.BadRequest(w,
			gecho.WithMessage("Invalid user ID"),
		).Send(); err != nil {
			logger.Error("send user ID error", gecho.Field("error", err))
		}
		return
	}

	// Find user
	user, exists := users[id]
	if !exists {
		if err := gecho.NotFound(w,
			gecho.WithMessage(fmt.Sprintf("User with ID %d not found", id)),
		).Send(); err != nil {
			logger.Error("send not found error", gecho.Field("error", err))
		}
		return
	}

	if err := gecho.Success(w,
		gecho.WithData(user),
	).Send(); err != nil {
		logger.Error("send user response", gecho.Field("error", err))
	}
}
