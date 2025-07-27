package repository

import (
	"database/sql"
	"fmt"
	"time"

	"chronotes-template/backend/internal/models"
	"chronotes-template/backend/pkg/database"
)

// UserRepository defines the interface for user database operations
type UserRepository interface {
	GetByID(id int64) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Create(user *models.CreateUserRequest) (*models.User, error)
	Update(id int64, user *models.UpdateUserRequest) (*models.User, error)
	Delete(id int64) error
	List(limit, offset int) ([]*models.User, error)
}

// userRepository implements UserRepository
type userRepository struct {
	db *database.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *database.DB) UserRepository {
	return &userRepository{db: db}
}

// GetByID retrieves a user by ID
func (r *userRepository) GetByID(id int64) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, email, name, created_at, updated_at FROM users WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return user, nil
}

// GetByEmail retrieves a user by email
func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, email, name, created_at, updated_at FROM users WHERE email = $1`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return user, nil
}

// Create creates a new user
func (r *userRepository) Create(req *models.CreateUserRequest) (*models.User, error) {
	now := time.Now()
	query := `INSERT INTO users (email, name, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`

	var id int64
	err := r.db.QueryRow(query, req.Email, req.Name, now, now).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &models.User{
		ID:        id,
		Email:     req.Email,
		Name:      req.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// Update updates an existing user
func (r *userRepository) Update(id int64, req *models.UpdateUserRequest) (*models.User, error) {
	// First, get the current user
	user, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update only provided fields
	if req.Name != nil {
		user.Name = *req.Name
	}

	user.UpdatedAt = time.Now()

	query := `UPDATE users SET name = $1, updated_at = $2 WHERE id = $3`
	_, err = r.db.Exec(query, user.Name, user.UpdatedAt, id)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

// Delete deletes a user by ID
func (r *userRepository) Delete(id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// List retrieves a list of users with pagination
func (r *userRepository) List(limit, offset int) ([]*models.User, error) {
	query := `SELECT id, email, name, created_at, updated_at FROM users LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user row: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return users, nil
}
