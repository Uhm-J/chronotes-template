package routes

import (
	"net/http"

	"chronotes-template/backend/internal/handlers"
	"chronotes-template/backend/internal/middleware"
	"chronotes-template/backend/internal/repository"

	"github.com/go-chi/chi/v5"
)

func SetupAuthRoutes(r chi.Router, h *handlers.AuthHandler, repo repository.UserRepository) {
	r.Route("/v1", func(r chi.Router) {
		r.Get("/auth/google/login", h.GoogleLogin)
		r.Get("/auth/google/callback", h.GoogleCallback)
		r.Post("/auth/logout", h.Logout)

		r.Group(func(r chi.Router) {
			r.Use(func(next http.Handler) http.Handler {
				return middleware.AuthMiddleware(repo)(next)
			})
			r.Get("/profile", h.GetCurrentUser)
		})
	})
}
