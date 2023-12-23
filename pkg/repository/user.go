package repository

import (
    "clinic-appointment-system/pkg/model"
    "gorm.io/gorm"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
    Create(user model.User) (model.User, error)
    FindByID(userID uint) (model.User, error)
    Update(user model.User) (model.User, error)
    Delete(userID uint) error
    FindByEmail(email string, user *model.User) error
}

type userRepository struct {
    db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{
        db: db,
    }
}

// Create adds a new user to the database
func (r *userRepository) Create(user model.User) (model.User, error) {
    err := r.db.Create(&user).Error
    return user, err
}

// FindByID finds a user by their ID
func (r *userRepository) FindByID(userID uint) (model.User, error) {
    var user model.User
    err := r.db.Where("id = ?", userID).First(&user).Error
    return user, err
}

// Update modifies an existing user in the database
func (r *userRepository) Update(user model.User) (model.User, error) {
    err := r.db.Save(&user).Error
    return user, err
}

// Delete removes a user from the database
func (r *userRepository) Delete(userID uint) error {
    err := r.db.Delete(&model.User{}, userID).Error
    return err
}

func (r *userRepository) FindByEmail(email string, user *model.User) error {
    result := r.db.Where("email = ?", email).First(user)
    return result.Error
}
