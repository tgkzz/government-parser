package service

import (
	"parser/internal/models"
	"parser/internal/repository"
	"parser/internal/service/samruk"
)

type Serivce struct {
	Samruk samruk.ISamrukService
	URLs   models.URL
}

func NewService(repo repository.IRepository, url models.URL) *Serivce {
	return &Serivce{
		Samruk: samruk.NewSamrukService(repo, url),
		URLs:   url,
	}
}
