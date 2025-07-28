package middleware

import (
	"context"
	"net/http"
	"strconv"

	"backend/internal/repository"
	"backend/pkg/response"
)

type ctxKey string

const UserKey ctxKey = "user"

func AuthMiddleware(repo repository.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session")
			if err != nil || cookie.Value == "" {
				response.Unauthorized(w, "Authentication required")
				return
			}
			id, err := strconv.ParseUint(cookie.Value, 10, 32)
			if err != nil {
				response.Unauthorized(w, "Invalid session")
				return
			}
			user, err := repo.GetByID(uint(id))
			if err != nil {
				response.Unauthorized(w, "User not found")
				return
			}
			ctx := context.WithValue(r.Context(), UserKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
