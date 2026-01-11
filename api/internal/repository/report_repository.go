package repository

import (
	"github.com/naotch/minibo/api/internal/model"
	"gorm.io/gorm"
)

type ReportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) FindTotal(userID uint) (int, error) {
	var total int
	err := r.db.Model(&model.Transaction{}).
		Where("user_id = ?", userID).
		Select("SUM(CASE WHEN category = 1 THEN amount ELSE -amount END)").
		Scan(&total).Error
	return total, err
}
