package service

import (
	"testing"

	"github.com/naotch/minibo/api/internal/model"
)

type mockTransactionRepository struct {
	transactions []*model.Transaction
}

func (m *mockTransactionRepository) CreateTransaction(t *model.Transaction) error {
	m.transactions = append(m.transactions, t)
	return nil
}

func TestTransactionService(t *testing.T) {

	mockRepo := &mockTransactionRepository{transactions: make([]*model.Transaction, 0)}
	service := NewTransactionService(mockRepo)

	userId := uint(1)
	title := "ランチ代"
	amount := 1200

	t.Run("Record_Expense", func(t *testing.T) {
		err := service.Record(userId, title, int(model.EXP), amount)
		if err != nil {
			t.Errorf("Record failed: %v", err)
		}

		lastTx := mockRepo.transactions[len(mockRepo.transactions)-1]
		if lastTx.Category != model.EXP {
			t.Errorf("Expected EXP, got %v", lastTx.Category)
		}
	})

	t.Run("Record_Income", func(t *testing.T) {
		err := service.Record(userId, "給料", int(model.INC), 300000)
		if err != nil {
			t.Errorf("Record failed: %v", err)
		}

		lastTx := mockRepo.transactions[len(mockRepo.transactions)-1]
		if lastTx.Category != model.INC {
			t.Errorf("Expected INC, got %v", lastTx.Category)
		}
	})

	t.Run("Record_InvalidCategory", func(t *testing.T) {
		invalidCat := 99
		err := service.Record(userId, title, invalidCat, amount)
		if err == nil {
			t.Error("Expected error for invalid category, but got nil")
		}
	})
}
