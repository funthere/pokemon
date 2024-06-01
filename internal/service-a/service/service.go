package service

import (
	"fmt"
	"math/rand"
	"time"
)

type SensorData struct {
	Value     float64
	Type      string
	ID1       string
	ID2       int
	Timestamp time.Time
}

var frequency int = 1000 // = 1 second

func UpdateFrequency(newFrequency int) {
	frequency = newFrequency
}

func GetFrequency() int {
	return frequency
}

func GenerateSensorData() {
	data := SensorData{
		Value:     rand.Float64() * 100,
		Type:      "temperature",
		ID1:       "SENSOR_A",
		ID2:       rand.Intn(100),
		Timestamp: time.Now(),
	}
	fmt.Println(fmt.Printf("%+v", data))
	// Send data to Microservice B
}
