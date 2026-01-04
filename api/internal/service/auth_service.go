package service

import (
	"github.com/naotch/minibo/api/internal/model"
	"github.com/naotch/minibo/api/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

type IAuthRepository interface {
	CreateUser(user *model.User) error
}

type AuthService struct {
	repository IAuthRepository
}

func NewAuthService(repository IAuthRepository) *AuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) Signup(email string, password string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Failed to hash password", err)
		return err
	}

	user := model.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.repository.CreateUser(&user)
	if err != nil {
		logger.Error("Failed to create user", err)
		return err
	}

	return nil
}
