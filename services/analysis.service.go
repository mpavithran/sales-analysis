package services

import (
	"github.com/mpavithran/sales-analysis/models"
	"github.com/mpavithran/sales-analysis/repositories"
)

type AnalysisService struct {
	repo *repositories.AnalysisRepository
}

func NewAnalysisService(repo *repositories.AnalysisRepository) *AnalysisService {
	return &AnalysisService{repo: repo}
}

func (s *AnalysisService) UploadCSV(filePath string) error {
	return s.repo.UploadCSV(filePath)
}

func (s *AnalysisService) GetRevenue(revenue models.Revenue) (float64, error) {
	return s.repo.GetRevenue(revenue)
}

func (s *AnalysisService) TopProducts(dateFrom string, dateTo string, n int) ([]models.TopProduct, error) {
	return s.repo.TopProducts(dateFrom, dateTo, n)
}
