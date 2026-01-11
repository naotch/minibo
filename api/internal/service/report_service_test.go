package service

import (
	"errors"
	"testing"
)

type mockReportRepository struct {
	total int
	err   error
}

func (m *mockReportRepository) FindTotal(userID uint) (int, error) {
	return m.total, m.err
}

func TestReportService(t *testing.T) {
	mockRepo := &mockReportRepository{}
	service := NewReportService(mockRepo)
	userID := uint(1)

	t.Run("ReportTotal_Success", func(t *testing.T) {
		mockRepo.total = 5000
		mockRepo.err = nil

		total, err := service.ReportTotal(userID)
		if err != nil {
			t.Errorf("ReportTotal failed: %v", err)
		}
		if total != 5000 {
			t.Errorf("Expected 5000, got %d", total)
		}
	})

	t.Run("ReportTotal_RepositoryError", func(t *testing.T) {
		mockRepo.err = errors.New("db error")

		_, err := service.ReportTotal(userID)
		if err == nil {
			t.Error("Expected error for repository failure, but got nil")
		}
	})
}
