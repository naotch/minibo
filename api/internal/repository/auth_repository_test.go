package repository

import (
	"testing"

	"github.com/naotch/minibo/api/internal/model"
)

func TestCreateUser(t *testing.T) {
	repository := NewAuthRepository(DB)
	testEmail := "test@example.com"
	testUser := &model.User{
		Email:    testEmail,
		Password: "hashedpassword",
	}

	t.Run("CreateUser", func(t *testing.T) {
		err := repository.CreateUser(testUser)

		if err != nil {
			t.Fatalf("Failed to create user: %v", err)
		}
	})

	t.Run("FindByEmail", func(t *testing.T) {
		user, err := repository.FindByEmail(testEmail)

		if err != nil {
			t.Fatalf("Failed to find user: %v", err)
		}

		if user.Email != testEmail {
			t.Errorf("Expected email %s, got %s", testEmail, user.Email)
		}
	})

	DB.Delete(testUser)
}
