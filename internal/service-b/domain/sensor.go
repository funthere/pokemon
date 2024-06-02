package domain

import "time"

type SensorData struct {
	ID          int       `json:"id"`
	SensorValue float64   `json:"value"`
	SensorType  string    `json:"type"`
	ID1         string    `json:"id1"`
	ID2         int       `json:"id2"`
	Timestamp   time.Time `json:"timestamp"`
}

type SensorRepository interface {
	Fetch(id1, id2, start, end string, limit, offset int) ([]SensorData, error)
	Delete(id1, id2, start, end string) (int64, error)
	Update(data SensorData, id1, id2, start, end string) (int64, error)
}
