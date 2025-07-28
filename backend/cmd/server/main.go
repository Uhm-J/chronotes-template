package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/gorm/logger"

	"chronotes-template/backend/config"
	"chronotes-template/backend/internal/handlers"
	"chronotes-template/backend/internal/models"
	"chronotes-template/backend/internal/repository"
	"chronotes-template/backend/internal/routes"
	"chronotes-template/backend/internal/services"
	"chronotes-template/backend/pkg/auth"
	"chronotes-template/backend/pkg/database"
)

func main() {
	cfg := config.Load()

	db, err := database.New(cfg.GetDatabaseURL(), logger.Info)
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer db.Close()

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	oauthService := auth.NewOAuthService(
		cfg.OAuth.GoogleClientID,
		cfg.OAuth.GoogleClientSecret,
		cfg.OAuth.GoogleRedirectURL,
	)
	authHandler := handlers.NewAuthHandler(userService, oauthService, cfg)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   cfg.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	routes.SetupAuthRoutes(r, authHandler, userRepo)

	r.Get("/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	if cfg.Server.FrontendPath != "" {
		if _, err := os.Stat(cfg.Server.FrontendPath); err == nil {
			fs := http.FileServer(http.Dir(cfg.Server.FrontendPath))
			r.Handle("/*", fs)
		}
	}

	log.Printf("starting server on :%s", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, r))
}
