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
