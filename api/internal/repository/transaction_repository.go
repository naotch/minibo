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

func (r *TransactionRepository) SummrizeTotal(userID uint) (int, error) {
	var total int
	err := r.db.Model(&model.Transaction{}).
		Where("user_id = ?", userID).
		Select("SUM(CASE WHEN category = 1 THEN amount ELSE -amount END)").
		Scan(&total).Error
	return total, err
}
