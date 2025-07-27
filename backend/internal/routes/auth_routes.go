package routes

import (
	"net/http"

	"chronotes-template/backend/internal/handlers"
	"chronotes-template/backend/internal/middleware"
	"chronotes-template/backend/internal/repository"

	"github.com/go-chi/chi/v5"
)

// SetupAuthRoutes sets up authentication routes
func SetupAuthRoutes(r chi.Router, authHandler *handlers.AuthHandler, userRepo repository.UserRepository) {
	r.Route("/v1/auth", func(r chi.Router) {
		// Public auth routes
		r.Get("/google/login", authHandler.GoogleLogin)
		r.Get("/google/callback", authHandler.GoogleCallback)

		// Routes that work with optional authentication
		r.With(func(next http.Handler) http.Handler {
			return middleware.OptionalAuthMiddleware(userRepo)(next)
		}).Get("/me", authHandler.GetCurrentUser)

		// Protected auth routes
		r.Group(func(r chi.Router) {
			r.Use(func(next http.Handler) http.Handler {
				return middleware.AuthMiddleware(userRepo)(next)
			})
			r.Post("/logout", authHandler.Logout)
		})
	})
}
