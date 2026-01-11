package service

import (
	"fmt"

	"github.com/naotch/minibo/api/internal/model"
	"github.com/naotch/minibo/api/pkg/logger"
)

type ITransactionRepository interface {
	CreateTransaction(transaction *model.Transaction) error
}

type TransactionService struct {
	repository ITransactionRepository
}

func NewTransactionService(repository ITransactionRepository) *TransactionService {
	return &TransactionService{repository: repository}
}

func (s *TransactionService) Record(userId uint, title string, category int, amount int) error {

	transactionCategory := model.TransactionCategory(category)
	if transactionCategory != model.EXP && transactionCategory != model.INC {
		return fmt.Errorf("invalid category: %d", category)
	}

	transaction := model.Transaction{
		UserID:   userId,
		Title:    title,
		Category: model.TransactionCategory(category),
		Amount:   amount,
	}

	err := s.repository.CreateTransaction(&transaction)
	if err != nil {
		logger.Error("Failed to create transaction", err)
		return err
	}

	return nil
}
