package service

type IReportRepository interface {
	FindTotal(userID uint) (int, error)
}

type ReportService struct {
	repository IReportRepository
}

func NewReportService(repository IReportRepository) *ReportService {
	return &ReportService{repository: repository}
}

func (s *ReportService) ReportTotal(userID uint) (int, error) {
	return s.repository.FindTotal(userID)
}
