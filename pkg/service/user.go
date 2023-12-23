package service

import (
    "clinic-appointment-system/pkg/model"
    "clinic-appointment-system/pkg/repository"
    "clinic-appointment-system/pkg/utils"
    "errors"
    "golang.org/x/crypto/bcrypt"
)

type UserService interface {
    CreateUser(user model.User) (model.User, error)
    GetUserByID(userID uint) (model.User, error)
    UpdateUser(user model.User) (model.User, error)
    DeleteUser(userID uint) error
    AuthenticateUser(email, password string) (model.User, error)
}

type userService struct {
    userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
    return &userService{
        userRepository: repo,
    }
}

func (s *userService) CreateUser(user model.User) (model.User, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return model.User{}, errors.New("failed to hash password")
    }
    user.Password = string(hashedPassword)

    createdUser, err := s.userRepository.Create(user)
    if err != nil {
        // Handle specific database errors if necessary
        return model.User{}, errors.New("failed to create user")
    }

    return createdUser, nil
}

func (s *userService) GetUserByID(userID uint) (model.User, error) {
    user, err := s.userRepository.FindByID(userID)
    if err != nil {
        return model.User{}, errors.New("user not found")
    }

    return user, nil
}

func (s *userService) UpdateUser(user model.User) (model.User, error) {
    // If password is being updated, hash the new password
    if user.Password != "" {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
        if err != nil {
            return model.User{}, errors.New("failed to hash password")
        }
        user.Password = string(hashedPassword)
    }

    updatedUser, err := s.userRepository.Update(user)
    if err != nil {
        return model.User{}, errors.New("failed to update user")
    }

    return updatedUser, nil
}

func (s *userService) DeleteUser(userID uint) error {
    err := s.userRepository.Delete(userID)
    if err != nil {
        return errors.New("failed to delete user")
    }
    return nil
}


// AuthenticateUser verifies a user's credentials and returns the user if they are valid
func (s *userService) AuthenticateUser(email, password string) (model.User, error) {
    var user model.User
    err := s.userRepository.FindByEmail(email, &user)
    if err != nil {
        return model.User{}, err
    }

    // Check if the hashed password matches the provided password
    if !utils.CheckPasswordHash(password, user.Password) {
        return model.User{}, errors.New("invalid credentials")
    }

    return user, nil
}

