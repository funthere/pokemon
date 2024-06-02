package domain

import (
	"math"
	"time"
)

type SensorData struct {
	ID          int       `json:"id"`
	SensorValue float64   `json:"value"`
	SensorType  string    `json:"type"`
	ID1         string    `json:"id1"`
	ID2         int       `json:"id2"`
	Timestamp   time.Time `json:"timestamp"`
}

type SensorRepository interface {
	Save(value float64, typ, id1 string, id2 int, timestamp string) error
	Fetch(id1, id2, start, end string, pagination *Pagination) ([]SensorData, error)
	Delete(id1, id2, start, end string) (int64, error)
	Update(data SensorDataUpdateReq, id1, id2, start, end string) (int64, error)
}

type SensorDataUpdateReq struct {
	SensorValue float64 `json:"value"`
}

type Pagination struct {
	Page       uint `json:"page"`
	Size       uint `json:"size"`
	TotalPages uint `json:"total_pages"`
	TotalRows  uint `json:"total_rows"`
}

func (p *Pagination) Init() {
	if p.Page == 0 {
		p.Page = 1
	}
	if p.Size == 0 {
		p.Size = 10
	}
}
func (p *Pagination) GetOffset() (offset uint) {
	offset = (p.Page - 1) * p.Size
	return
}
func (p *Pagination) CalculateTotalPages() {
	if p.Size == 0 {
		p.TotalPages = 0
	} else {
		p.TotalPages = uint(math.Ceil(float64(p.TotalRows) / float64(p.Size)))
	}
}
