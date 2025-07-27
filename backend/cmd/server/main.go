package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"chronotes-template/backend/config"
	"chronotes-template/backend/internal/handlers"
	"chronotes-template/backend/internal/repository"
	"chronotes-template/backend/internal/routes"
	"chronotes-template/backend/internal/services"
	"chronotes-template/backend/pkg/auth"
	"chronotes-template/backend/pkg/database"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := database.New(cfg.GetDatabaseURL())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Run migrations
	if err := db.Migrate(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)

	// Initialize OAuth service
	oauthService := auth.NewOAuthService(
		cfg.OAuth.GoogleClientID,
		cfg.OAuth.GoogleClientSecret,
		cfg.OAuth.GoogleRedirectURL,
	)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(userService, oauthService, cfg)

	// Initialize Chi router
	r := chi.NewRouter()

	// Setup global middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	// Setup CORS
	corsHandler := cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:3000"}, // Frontend URLs
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(corsHandler)

	// Setup API routes
	routes.SetupAuthRoutes(r, authHandler, userRepo)
	routes.SetupUserRoutes(r, authHandler, userRepo)

	// API health check
	r.Get("/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok", "service": "chronotes-backend"}`))
	})

	// Serve frontend static files
	if _, err := os.Stat(cfg.Server.FrontendPath); err == nil {
		log.Printf("Serving frontend from: %s", cfg.Server.FrontendPath)
		fs := http.FileServer(http.Dir(cfg.Server.FrontendPath))
		r.Handle("/*", fs)
	} else {
		log.Printf("Frontend not found at: %s", cfg.Server.FrontendPath)
		r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Frontend not built – run `npm run build` in the frontend directory"))
		})
	}

	// Start server
	log.Printf("🚀 Starting server on port %s", cfg.Server.Port)
	log.Printf("📝 Environment: %s", cfg.Server.Environment)
	log.Printf("🗄️  Database: PostgreSQL at %s:%d", cfg.Database.Host, cfg.Database.Port)
	log.Printf("🌐 API Base URL: http://localhost:%s/v1", cfg.Server.Port)

	if err := http.ListenAndServe(":"+cfg.Server.Port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
