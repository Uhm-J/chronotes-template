package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	OAuth    OAuthConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port         string
	Environment  string
	FrontendPath string
	FrontendURL  string
}

// DatabaseConfig holds database-related configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

// OAuthConfig holds OAuth-related configuration
type OAuthConfig struct {
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURL  string
}

// Load loads configuration from environment variables and .env file
func Load() *Config {
	// Try to load .env file (ignore error if file doesn't exist)
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, using environment variables: %v", err)
	}

	config := &Config{
		Server: ServerConfig{
			Port:         getEnv("PORT", "8080"),
			Environment:  getEnv("ENVIRONMENT", "development"),
			FrontendPath: getEnv("FRONTEND_PATH", "frontend/dist"),
			FrontendURL:  getEnv("FRONTEND_URL", "http://localhost:2010"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "chronotes"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		OAuth: OAuthConfig{
			GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
			GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
			GoogleRedirectURL:  getEnv("GOOGLE_REDIRECT_URL", ""),
		},
	}

	// Validate required configuration
	if err := config.Validate(); err != nil {
		log.Printf("Configuration warning: %v", err)
	}

	return config
}

// LoadFromFile loads configuration from a specific .env file
func LoadFromFile(filename string) *Config {
	if err := godotenv.Load(filename); err != nil {
		log.Fatalf("Error loading %s file: %v", filename, err)
	}
	return Load()
}

// Validate checks if required configuration is present
func (c *Config) Validate() error {
	if c.OAuth.GoogleClientID == "" || c.OAuth.GoogleClientSecret == "" || c.OAuth.GoogleRedirectURL == "" {
		log.Println("Warning: Google OAuth credentials not set. Login will fail.")
	}

	if c.Database.Name == "" {
		log.Println("Warning: Database name not set.")
	}

	if c.Database.Password == "" && c.IsDevelopment() {
		log.Println("Warning: Database password not set.")
	}

	return nil
}

// IsDevelopment returns true if the environment is development
func (c *Config) IsDevelopment() bool {
	return c.Server.Environment == "development"
}

// IsProduction returns true if the environment is production
func (c *Config) IsProduction() bool {
	return c.Server.Environment == "production"
}

// GetDatabaseURL returns the PostgreSQL connection string
func (c *Config) GetDatabaseURL() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
		c.Database.SSLMode,
	)
}

// getEnv gets an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// getEnvAsInt gets an environment variable as integer with a fallback value
func getEnvAsInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}

// getEnvAsBool gets an environment variable as boolean with a fallback value
func getEnvAsBool(key string, fallback bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return fallback
}
