package handlers

import (
	"fmt"
	"net/http"

	"chronotes-template/backend/config"
	"chronotes-template/backend/internal/middleware"
	"chronotes-template/backend/internal/models"
	"chronotes-template/backend/internal/services"
	"chronotes-template/backend/pkg/auth"
	"chronotes-template/backend/pkg/response"
)

type AuthHandler struct {
	userService  services.UserService
	oauthService *auth.OAuthService
	cfg          *config.Config
}

func NewAuthHandler(us services.UserService, oauth *auth.OAuthService, cfg *config.Config) *AuthHandler {
	return &AuthHandler{userService: us, oauthService: oauth, cfg: cfg}
}

func (h *AuthHandler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := h.oauthService.GetAuthURL("state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *AuthHandler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		response.BadRequest(w, "Authorization code missing")
		return
	}

	token, err := h.oauthService.ExchangeCode(r.Context(), code)
	if err != nil {
		response.InternalServerError(w, "Token exchange failed")
		return
	}

	info, err := h.oauthService.GetUserInfo(r.Context(), token)
	if err != nil {
		response.InternalServerError(w, "Failed to fetch user info")
		return
	}

	user, err := h.userService.GetOrCreateFromOAuth(info.Email, info.Name)
	if err != nil {
		response.InternalServerError(w, "User creation failed")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    fmt.Sprintf("%d", user.ID),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   h.cfg.CookieSecure,
	})

	http.Redirect(w, r, h.cfg.Server.FrontendURL, http.StatusTemporaryRedirect)
}

// Logout clears the session cookie
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   h.cfg.CookieSecure,
	})
	response.Success(w, map[string]string{"message": "logged out"})
}

func (h *AuthHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	u := r.Context().Value(middleware.UserKey)
	if u == nil {
		response.Unauthorized(w, "Unauthorized")
		return
	}
	user, ok := u.(*models.User)
	if !ok {
		response.InternalServerError(w, "context error")
		return
	}
	response.Success(w, user.ToResponse())
}
