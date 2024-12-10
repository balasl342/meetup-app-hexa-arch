package user

import (
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
