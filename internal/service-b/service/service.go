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

func (s *Service) GetData(id1, id2, start, end, page, size string) ([]repository.SensorData, error) {
	return s.repo.FindByIDAndTime(id1, id2, start, end, page, size)
}

func (s *Service) DeleteData(id1, id2, start, end string) error {
	return s.repo.DeleteByIDAndTime(id1, id2, start, end)
}

func (s *Service) UpdateData(value float64, typ, id1 string, id2 int, timestamp string) error {
	return s.repo.Update(value, typ, id1, id2, timestamp)
}

func (s *Service) GetDataByIDCombination(id1, id2 string) ([]repository.SensorData, error) {
	return s.repo.FindByIDCombination(id1, id2)
}

func (s *Service) GetDataByDuration(start, end string) ([]repository.SensorData, error) {
	return s.repo.FindByDuration(start, end)
}

func (s *Service) GetDataByIDAndTimestamp(id1, id2, start, end string) ([]repository.SensorData, error) {
	return s.repo.FindByIDAndTimestamp(id1, id2, start, end)
}
