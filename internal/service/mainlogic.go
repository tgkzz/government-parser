package service

import (
	"log"
	"parser/internal/models"
)

func (s *Serivce) Execute() {
	filter := models.Filter{
		SizeOfParsing: 8,
		DateFrom:      "2022-01-01T00:00:00+06:00",
		DateTo:        "2022-02-01T00:00:00+06:00",
	}

	for i := 0; i <= filter.SizeOfParsing; i++ {
		if err := s.Samruk.UploadLotsByFilter(filter, i); err != nil {
			log.Print(err)
			return
		}
	}
}
