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
	UploadLotsByFilter(filter models.Filter, id int) error
	GetLotsById(id string) (models.Result, error)
}

func NewSamrukService(repo repository.IRepository, URLs models.URL) *SamrukService {
	return &SamrukService{
		repo: repo,
		URLs: URLs,
	}
}
