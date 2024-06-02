package repository

import (
	"database/sql"

	"github.com/funthere/pokemon/internal/service-b/domain"
)

type mysqlSensorRepository struct {
	Conn *sql.DB
}

func NewMysqlSensorRepository(Conn *sql.DB) domain.SensorRepository {
	return &mysqlSensorRepository{Conn}
}

func (m *mysqlSensorRepository) Fetch(id1, id2, start, end string, limit, offset int) ([]domain.SensorData, error) {
	query := "SELECT id, value, type, id1, id2, timestamp FROM sensor_data WHERE 1=1"
	args := make([]any, 0)

	if id1 != "" {
		query += " AND id1 = ?"
		args = append(args, id1)
	}
	if id2 != "" {
		query += " AND id2 = ?"
		args = append(args, id2)
	}
	if start != "" && end != "" {
		query += " AND timestamp BETWEEN ? AND ?"
		args = append(args, start, end)
	}
	args = append(args, limit, offset)
	query += " LIMIT ? OFFSET ?"

	// fmt.Println("===", query, args)
	rows, err := m.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make([]domain.SensorData, 0)
	for rows.Next() {
		var sd domain.SensorData
		if err := rows.Scan(&sd.ID, &sd.SensorValue, &sd.SensorType, &sd.ID1, &sd.ID2, &sd.Timestamp); err != nil {
			return nil, err
		}
		data = append(data, sd)
	}

	return data, nil
}

func (m *mysqlSensorRepository) Delete(id1, id2, start, end string) (int64, error) {
	query := "DELETE FROM sensor_data WHERE 1=1"
	args := make([]any, 0)

	if id1 != "" {
		query += " AND id1 = ?"
		args = append(args, id1)
	}
	if id2 != "" {
		query += " AND id2 = ?"
		args = append(args, id2)
	}
	if start != "" && end != "" {
		query += " AND timestamp BETWEEN ? AND ?"
		args = append(args, start, end)
	}

	res, err := m.Conn.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (m *mysqlSensorRepository) Update(data domain.SensorData, id1, id2, start, end string) (int64, error) {
	query := "UPDATE sensor_data SET value = ? WHERE 1=1"
	args := []any{data.SensorValue}

	if id1 != "" {
		query += " AND id1 = ?"
		args = append(args, id1)
	}
	if id2 != "" {
		query += " AND id2 = ?"
		args = append(args, id2)
	}
	if start != "" && end != "" {
		query += " AND timestamp BETWEEN ? AND ?"
		args = append(args, start, end)
	}

	res, err := m.Conn.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}