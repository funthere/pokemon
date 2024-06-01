package service

import (
	"database/sql"

	"github.com/funthere/pokemon/internal/service-b/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(db *sql.DB) *Service {
	return &Service{
		repo: repository.NewRepository(db),
	}
}

func (s *Service) Fetch(page, size uint32) ([]repository.SensorData, error) {
	return s.repo.Fetch(page, size)
}

func (s *Service) SaveData(value float64, typ, id1 string, id2 int, timestamp string) error {
	return s.repo.Save(value, typ, id1, id2, timestamp)
}
