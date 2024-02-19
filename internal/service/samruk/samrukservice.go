package samruk

import (
	"parser/internal/models"
	"parser/internal/repository"
)

type SamrukService struct {
	repo repository.IRepository
	URLs models.URL
}

type ISamrukService interface {
	GetLotsByFilter(filter models.Filter) ([]models.Result, error)
	GetLotsById(id string) (models.Result, error)
}

func NewSamrukService(repo repository.IRepository, URLs models.URL) *SamrukService {
	return &SamrukService{
		repo: repo,
		URLs: URLs,
	}
}
