package services

import (
	"fmt"

	"backend/internal/models"
	"backend/internal/repository"

	"gorm.io/gorm"
)

type UserService interface {
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	CreateUser(email, name string) (*models.User, error)
	GetOrCreateFromOAuth(email, name string) (*models.User, error)
}

type userService struct{ repo repository.UserRepository }

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetByID(id uint) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) GetByEmail(email string) (*models.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *userService) CreateUser(email, name string) (*models.User, error) {
	if email == "" {
		return nil, fmt.Errorf("email is required")
	}
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}

	if existing, err := s.repo.GetByEmail(email); err == nil && existing != nil {
		return nil, fmt.Errorf("user with email %s already exists", email)
	}

	user := &models.User{Email: email, Name: name}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetOrCreateFromOAuth(email, name string) (*models.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return s.CreateUser(email, name)
		}
		return nil, err
	}
	return user, nil
}
