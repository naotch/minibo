package repository

import (
	"github.com/naotch/minibo/api/internal/model"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransaction(transaction *model.Transaction) error {
	if err := r.db.Create(transaction).Error; err != nil {
		return err
	}
	return nil
}
