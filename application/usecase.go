package application

import (
	"errors"
	"hexagonal_both_publicidad_go/domain"
)

// Servicio de aplicación que orquesta scraping y almacenamiento
type ScrapeService struct {
	scraper domain.Scraper
	repo    domain.DataRepository
}

func NewScrapeService(scraper domain.Scraper, repo domain.DataRepository) *ScrapeService {
	return &ScrapeService{
		scraper: scraper,
		repo:    repo,
	}
}

func (s *ScrapeService) ScrapeAndSave(url string) error {
	// Realizar scraping
	data, err := s.scraper.Scrape(url)
	if err != nil {
		return err
	}

	// Verificar y filtrar duplicados
	var uniqueData []domain.Data
	for _, item := range data {
		isDup, err := s.repo.IsDuplicate(item.Link)
		if err != nil {
			return err
		}
		if !isDup {
			uniqueData = append(uniqueData, item)
		}
	}

	// Guardar datos no duplicados
	if len(uniqueData) == 0 {
		return errors.New("no se encontraron datos únicos para guardar")
	}

	return s.repo.Save(uniqueData)
}
