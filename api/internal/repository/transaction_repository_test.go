package repository

import (
	"testing"

	"github.com/naotch/minibo/api/internal/model"
)

func TestCreateTransaction(t *testing.T) {
	repository := NewTransactionRepository(DB)
	testTran := &model.Transaction{
		UserID:   1,
		Title:    "test title",
		Category: 0,
		Amount:   10000,
	}

	t.Run("CreateTransaction", func(t *testing.T) {
		err := repository.CreateTransaction(testTran)
		if err != nil {
			t.Fatalf("Failed to create transaction: %v", err)
		}
	})
}
