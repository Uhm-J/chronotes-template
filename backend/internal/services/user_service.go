package services

import (
	"fmt"

	"chronotes-template/backend/internal/models"
	"chronotes-template/backend/internal/repository"
)

// UserService defines the interface for user business logic
type UserService interface {
	GetByID(id int64) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	CreateUser(email, name string) (*models.User, error)
	UpdateUser(id int64, req *models.UpdateUserRequest) (*models.User, error)
	DeleteUser(id int64) error
	ListUsers(page, limit int) ([]*models.User, error)
	CreateOrUpdateFromOAuth(email, name string) (*models.User, error)
}

// userService implements UserService
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// GetByID retrieves a user by ID
func (s *userService) GetByID(id int64) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

// GetByEmail retrieves a user by email
func (s *userService) GetByEmail(email string) (*models.User, error) {
	return s.userRepo.GetByEmail(email)
}

// CreateUser creates a new user with validation
func (s *userService) CreateUser(email, name string) (*models.User, error) {
	// Validate input
	if email == "" {
		return nil, fmt.Errorf("email is required")
	}
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}

	// Check if user already exists
	existingUser, err := s.userRepo.GetByEmail(email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", email)
	}

	// Create user
	req := &models.CreateUserRequest{
		Email: email,
		Name:  name,
	}

	return s.userRepo.Create(req)
}

// UpdateUser updates an existing user
func (s *userService) UpdateUser(id int64, req *models.UpdateUserRequest) (*models.User, error) {
	// Validate that user exists
	_, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return s.userRepo.Update(id, req)
}

// DeleteUser deletes a user
func (s *userService) DeleteUser(id int64) error {
	// Validate that user exists
	_, err := s.userRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	return s.userRepo.Delete(id)
}

// ListUsers lists users with pagination
func (s *userService) ListUsers(page, limit int) ([]*models.User, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit
	return s.userRepo.List(limit, offset)
}

// CreateOrUpdateFromOAuth creates a new user or updates existing one from OAuth data
func (s *userService) CreateOrUpdateFromOAuth(email, name string) (*models.User, error) {
	// Try to get existing user
	existingUser, err := s.userRepo.GetByEmail(email)
	if err != nil {
		// User doesn't exist, create new one
		return s.CreateUser(email, name)
	}

	// User exists, update name if different
	if existingUser.Name != name {
		updateReq := &models.UpdateUserRequest{
			Name: &name,
		}
		return s.userRepo.Update(existingUser.ID, updateReq)
	}

	return existingUser, nil
} 