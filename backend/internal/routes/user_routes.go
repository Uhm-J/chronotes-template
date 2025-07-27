package routes

import (
	"net/http"

	"chronotes-template/backend/internal/handlers"
	"chronotes-template/backend/internal/middleware"
	"chronotes-template/backend/internal/repository"

	"github.com/go-chi/chi/v5"
)

// SetupUserRoutes sets up user-related routes
func SetupUserRoutes(r chi.Router, authHandler *handlers.AuthHandler, userRepo repository.UserRepository) {
	r.Route("/v1/users", func(r chi.Router) {
		// All user routes require authentication
		r.Use(func(next http.Handler) http.Handler {
			return middleware.AuthMiddleware(userRepo)(next)
		})

		// GET /v1/users/{id} - Get user by ID
		r.Get("/{id}", authHandler.GetUserByID)

		// Future user management endpoints can be added here:
		// r.Get("/", userHandler.ListUsers)        // List users with pagination
		// r.Put("/{id}", userHandler.UpdateUser)   // Update user
		// r.Delete("/{id}", userHandler.DeleteUser) // Delete user
	})
}
