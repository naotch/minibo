package repository

import (
	"testing"

	"github.com/naotch/minibo/api/internal/model"
)

func TestFindTotal(t *testing.T) {
	repository := NewReportRepository(DB)
	userID := uint(100)

	t.Cleanup(func() {
		DB.Unscoped().Where("user_id = ?", userID).Delete(&model.Transaction{})
	})

	t.Run("FindTotal", func(t *testing.T) {
		DB.Create(&model.Transaction{UserID: userID, Title: "臨時収入", Category: model.INC, Amount: 5000})
		DB.Create(&model.Transaction{UserID: userID, Title: "飲み代", Category: model.EXP, Amount: 2000})
		total, err := repository.FindTotal(userID)
		if err != nil {
			t.Fatalf("Failed to find total: %v", err)
		}
		if total < 3000 {
			t.Errorf("Expected at least 3000, but got %d", total)
		}
	})
}
