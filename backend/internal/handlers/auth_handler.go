package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"chronotes-template/backend/config"
	"chronotes-template/backend/internal/middleware"
	"chronotes-template/backend/internal/models"
	"chronotes-template/backend/internal/services"
	"chronotes-template/backend/pkg/auth"
	"chronotes-template/backend/pkg/response"
)

// AuthHandler handles authentication-related HTTP requests
type AuthHandler struct {
	userService  services.UserService
	oauthService *auth.OAuthService
	config       *config.Config
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(userService services.UserService, oauthService *auth.OAuthService, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		userService:  userService,
		oauthService: oauthService,
		config:       cfg,
	}
}

// GoogleLogin initiates the Google OAuth flow
func (h *AuthHandler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := h.oauthService.GetAuthURL("state-token") // In production, use a proper state token
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// GoogleCallback handles the Google OAuth callback
func (h *AuthHandler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	// Get authorization code
	code := r.FormValue("code")
	if code == "" {
		response.BadRequest(w, "Authorization code not provided")
		return
	}

	// Exchange code for token
	token, err := h.oauthService.ExchangeCode(r.Context(), code)
	if err != nil {
		response.InternalServerError(w, "Failed to exchange authorization code")
		return
	}

	// Get user info from Google
	userInfo, err := h.oauthService.GetUserInfo(r.Context(), token)
	if err != nil {
		response.InternalServerError(w, "Failed to get user information")
		return
	}

	// Create or update user in database
	user, err := h.userService.CreateOrUpdateFromOAuth(userInfo.Email, userInfo.Name)
	if err != nil {
		response.InternalServerError(w, "Failed to create or update user")
		return
	}

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    fmt.Sprintf("%d", user.ID),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		// In production, set Secure: true when using HTTPS
	})

	// Redirect to frontend
	http.Redirect(w, r, h.config.Server.FrontendURL, http.StatusTemporaryRedirect)
}

// GetCurrentUser returns the current authenticated user
func (h *AuthHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	// Get user from context (set by auth middleware)
	userCtx := r.Context().Value(middleware.UserKey)
	if userCtx == nil {
		response.Unauthorized(w, "User not found in context")
		return
	}

	user, ok := userCtx.(*models.User)
	if !ok {
		response.InternalServerError(w, "Invalid user type in context")
		return
	}

	response.Success(w, user.ToResponse())
}

// Logout logs out the current user
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Clear session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	response.Success(w, map[string]string{"message": "Logged out successfully"})
}

// GetUserByID returns a user by ID (protected route) using Chi URL parameters
func (h *AuthHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Get user ID from Chi URL parameter
	idStr := r.URL.Query().Get("id") // Changed from chi.URLParam to r.URL.Query().Get
	if idStr == "" {
		response.BadRequest(w, "User ID is required")
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(w, "Invalid user ID")
		return
	}

	user, err := h.userService.GetByID(id)
	if err != nil {
		response.NotFound(w, "User not found")
		return
	}

	response.Success(w, user.ToResponse())
}
