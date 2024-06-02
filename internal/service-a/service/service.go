package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	pb "github.com/funthere/pokemon/proto"
)

type SensorData struct {
	Value     float64
	Type      string
	ID1       string
	ID2       int
	Timestamp time.Time
}

var frequency int = 1 // = 1 second

func UpdateFrequency(newFrequency int) {
	frequency = newFrequency
}

func GetFrequency() int {
	return frequency
}

func GetType() string {
	if os.Getenv("SENSOR_TYPE") != "" {
		return os.Getenv("SENSOR_TYPE")
	}
	return "temperature"
}

func GenerateData(client pb.SensorServiceClient) {
	for {
		data := &pb.SensorData{
			Value:     rand.Float32() * 100,
			Type:      GetType(),
			Id1:       string(randomUppercaseLetter()),
			Id2:       int32(rand.Intn(100)),
			Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		}

		res, err := client.SendSensorData(context.Background(), data)
		if err != nil {
			log.Printf("could not send data: %v", err)
		}

		fmt.Printf("%+v ==> %s\n", data, res.GetStatus())

		sleep := 1000 / GetFrequency() // Frequency per second
		time.Sleep(time.Millisecond * time.Duration(sleep))
	}
}
func randomUppercaseLetter() rune {
	return rune('A' + rand.Intn('Z'-'A'+1))
}
