package repository

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

type SensorData struct {
	Value     float64 `json:"value"`
	Type      string  `json:"type"`
	ID1       string  `json:"id1"`
	ID2       int     `json:"id2"`
	Timestamp string  `json:"timestamp"`
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Fetch(page, size uint32) ([]SensorData, error) {
	offset := uint32(0)
	if size == 0 {
		size = 10
	}
	if page == 0 {
		page = 1
	} else {
		offset = size * (page - 1)
	}

	query := "SELECT value, type, id1, id2, timestamp FROM sensor_data LIMIT ? OFFSET ?"
	rows, err := r.db.Query(query, size, offset)
	fmt.Println(query, size, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []SensorData
	for rows.Next() {
		var d SensorData
		if err := rows.Scan(&d.Value, &d.Type, &d.ID1, &d.ID2, &d.Timestamp); err != nil {
			return nil, err
		}
		data = append(data, d)
	}
	return data, nil
}
