package service

import (
	"os"
	"testing"

	"github.com/naotch/minibo/api/internal/model"
	"gorm.io/gorm"
)

type mockAuthRepository struct {
	users map[string]*model.User
}

func (m *mockAuthRepository) CreateUser(user *model.User) error {
	m.users[user.Email] = user
	return nil
}

func (m *mockAuthRepository) FindByEmail(email string) (*model.User, error) {
	user, ok := m.users[email]
	if !ok {
		return nil, gorm.ErrRecordNotFound
	}
	return user, nil
}

func TestAuthService(t *testing.T) {
	os.Setenv("SECRET_KEY", "test-secret")
	mockRepo := &mockAuthRepository{users: make(map[string]*model.User)}
	service := NewAuthService(mockRepo)

	email := "test@example.com"
	password := "password123"

	t.Run("Signup", func(t *testing.T) {
		err := service.Signup(email, password)
		if err != nil {
			t.Errorf("Signup failed: %v", err)
		}
	})

	t.Run("Signin", func(t *testing.T) {
		token, err := service.Signin(email, password)
		if err != nil {
			t.Errorf("Signin failed: %v", err)
		}
		if token == "" {
			t.Error("Token should not be empty")
		}
	})

	t.Run("Signin_InvalidPassword", func(t *testing.T) {
		_, err := service.Signin(email, "wrongpassword")
		if err == nil {
			t.Error("Expected error for wrong password, but got nil")
		}
	})
}
