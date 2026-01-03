package user

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id" example:"1"`
	Name      string         `gorm:"size:100;not null" json:"name" example:"John Doe"`
	Email     string         `gorm:"size:100;uniqueIndex;not null" json:"email" example:"john@example.com"`
	Password  string         `gorm:"size:255;not null" json:"-"` // "-" hides from JSON
	CreatedAt time.Time      `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2024-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete
}

// TableName overrides the table name
func (User) TableName() string {
	return "users"
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100" example:"John Doe"`
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"secret123"`
}

// UpdateUserRequest represents the request body for updating a user
type UpdateUserRequest struct {
	Name  string `json:"name" binding:"omitempty,min=2,max=100" example:"John Updated"`
	Email string `json:"email" binding:"omitempty,email" example:"john.updated@example.com"`
}

// UserResponse represents the response body for user data
type UserResponse struct {
	ID        uint      `json:"id" example:"1"`
	Name      string    `json:"name" example:"John Doe"`
	Email     string    `json:"email" example:"john@example.com"`
	CreatedAt time.Time `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}

// ToResponse converts User to UserResponse
func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// LoginRequest represents the request body for user login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required" example:"secret123"`
}

// LoginResponse represents the response body for login
type LoginResponse struct {
	Token string        `json:"token" example:"eyJhbGciOiJIUzI1NiIs..."`
	User  *UserResponse `json:"user"`
}
