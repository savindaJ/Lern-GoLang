package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// UserService interface defines the contract for user business logic
type UserService interface {
	Register(req *CreateUserRequest) (*UserResponse, error)
	Login(req *LoginRequest) (*User, error)
	GetByID(id uint) (*UserResponse, error)
	GetAll(page, limit int) ([]UserResponse, int64, error)
	Update(id uint, req *UpdateUserRequest) (*UserResponse, error)
	Delete(id uint) error
}

// userService implements UserService
type userService struct {
	repo UserRepository
}

// NewUserService creates a new user service
func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

// Register creates a new user
func (s *userService) Register(req *CreateUserRequest) (*UserResponse, error) {
	// Check if email already exists
	existingUser, err := s.repo.FindByEmail(req.Email)
	if err == nil && existingUser != nil {
		return nil, ErrEmailAlreadyExists
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user.ToResponse(), nil
}

// Login authenticates a user
func (s *userService) Login(req *LoginRequest) (*User, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

// GetByID retrieves a user by ID
func (s *userService) GetByID(id uint) (*UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user.ToResponse(), nil
}

// GetAll retrieves all users with pagination
func (s *userService) GetAll(page, limit int) ([]UserResponse, int64, error) {
	users, total, err := s.repo.FindAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]UserResponse, len(users))
	for i, user := range users {
		responses[i] = *user.ToResponse()
	}

	return responses, total, nil
}

// Update updates a user
func (s *userService) Update(id uint, req *UpdateUserRequest) (*UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// Check if new email already exists
	if req.Email != "" && req.Email != user.Email {
		existingUser, err := s.repo.FindByEmail(req.Email)
		if err == nil && existingUser != nil {
			return nil, ErrEmailAlreadyExists
		}
		user.Email = req.Email
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return user.ToResponse(), nil
}

// Delete deletes a user
func (s *userService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return err
	}
	return s.repo.Delete(id)
}
