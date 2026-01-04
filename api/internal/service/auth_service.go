package service

import (
	"errors"

	"github.com/naotch/minibo/api/internal/model"
	"github.com/naotch/minibo/api/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	CreateUser(user *model.User) error
	FindByEmail(email string) (*model.User, error)
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

func (s *AuthService) Signin(email string, password string) (string, error) {

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error("User not found", err)
			return "", errors.New("Invalid email or password")
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		logger.Error("Invalid email or password", err)
		return "", errors.New("invalid email or password")
	}
	return "dummy-token", nil
}
