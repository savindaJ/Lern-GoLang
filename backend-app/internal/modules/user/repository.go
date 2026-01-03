package user

import (
	"gorm.io/gorm"
)

// UserRepository interface defines the contract for user data access
type UserRepository interface {
	Create(user *User) error
	FindByID(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	FindAll(page, limit int) ([]User, int64, error)
	Update(user *User) error
	Delete(id uint) error
}

// userRepository implements UserRepository using GORM
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create creates a new user in the database
func (r *userRepository) Create(user *User) error {
	return r.db.Create(user).Error
}

// FindByID finds a user by ID
func (r *userRepository) FindByID(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail finds a user by email
func (r *userRepository) FindByEmail(email string) (*User, error) {
	var user User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindAll retrieves all users with pagination
func (r *userRepository) FindAll(page, limit int) ([]User, int64, error) {
	var users []User
	var total int64

	offset := (page - 1) * limit

	// Get total count
	if err := r.db.Model(&User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated users
	if err := r.db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// Update updates an existing user
func (r *userRepository) Update(user *User) error {
	return r.db.Save(user).Error
}

// Delete soft deletes a user by ID
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}
