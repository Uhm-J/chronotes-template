package middleware

import (
	"context"
	"net/http"
	"strconv"

	"chronotes-template/backend/internal/repository"
	"chronotes-template/backend/pkg/response"
)

// ChiAuthMiddleware creates Chi-compatible authentication middleware
func ChiAuthMiddleware(userRepo repository.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get session cookie
			cookie, err := r.Cookie("session")
			if err != nil || cookie.Value == "" {
				response.Unauthorized(w, "Authentication required")
				return
			}

			// Parse user ID from cookie
			userID, err := strconv.ParseInt(cookie.Value, 10, 64)
			if err != nil {
				response.Unauthorized(w, "Invalid session")
				return
			}

			// Get user from database
			user, err := userRepo.GetByID(userID)
			if err != nil {
				response.Unauthorized(w, "User not found")
				return
			}

			// Add user to context
			ctx := context.WithValue(r.Context(), UserKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// ChiOptionalAuthMiddleware creates Chi-compatible optional authentication middleware
func ChiOptionalAuthMiddleware(userRepo repository.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get session cookie
			cookie, err := r.Cookie("session")
			if err != nil || cookie.Value == "" {
				next.ServeHTTP(w, r)
				return
			}

			// Parse user ID from cookie
			userID, err := strconv.ParseInt(cookie.Value, 10, 64)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// Get user from database
			user, err := userRepo.GetByID(userID)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// Add user to context
			ctx := context.WithValue(r.Context(), UserKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
