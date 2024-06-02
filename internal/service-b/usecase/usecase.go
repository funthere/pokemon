package usecase

import (
	"github.com/funthere/pokemon/internal/service-b/domain"
)

type SensorUsecase interface {
	Fetch(id1, id2, start, end string, page, pageSize int) ([]domain.SensorData, error)
	Delete(id1, id2, start, end string) (int64, error)
	Update(data domain.SensorData, id1, id2, start, end string) (int64, error)
}

type sensorUsecase struct {
	sensorRepo domain.SensorRepository
}

func NewSensorUsecase(repo domain.SensorRepository) SensorUsecase {
	return &sensorUsecase{
		sensorRepo: repo,
	}
}

func (s *sensorUsecase) Fetch(id1, id2, start, end string, page, pageSize int) ([]domain.SensorData, error) {
	offset := (page - 1) * pageSize
	return s.sensorRepo.Fetch(id1, id2, start, end, pageSize, offset)
}

func (s *sensorUsecase) Delete(id1, id2, start, end string) (int64, error) {
	return s.sensorRepo.Delete(id1, id2, start, end)
}

func (s *sensorUsecase) Update(data domain.SensorData, id1, id2, start, end string) (int64, error) {
	return s.sensorRepo.Update(data, id1, id2, start, end)
}
