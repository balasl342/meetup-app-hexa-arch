package user

import (
	"context"
	"errors"
	"fmt"
	"meetup-app-hexa-arch/internal/adapters/auth"
)

type UserService struct {
	repo        *MongoDBUserRepository
	authService *auth.JWTService
}

func NewUserService(repo *MongoDBUserRepository, authService *auth.JWTService) *UserService {
	return &UserService{
		repo:        repo,
		authService: authService,
	}
}

// Register handles user registration.
func (u *UserService) Register(user User) error {
	// Example: Validate user input (expand as needed)
	if user.Email == "" || user.Password == "" {
		return errors.New("email and password are required")
	}

	// Save the user to MongoDB
	return u.repo.Save(context.Background(), &user)
}

// Login handles user login.
func (u *UserService) Login(email, password string) (*User, error) {
	// Retrieve the user from MongoDB
	user, err := u.repo.GetByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// For now, assume the password matches (expand to use hashed passwords)
	if user.Password != password {
		return nil, errors.New("invalid password")
	}

	// Generate JWT token using authService
	token, err := u.authService.GenerateToken(user.ID.Hex())
	if err != nil {
		return nil, err
	}
	fmt.Println("token", token)
	return user, nil
}
