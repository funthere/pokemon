package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type SensorData struct {
	Value     float64   `json:"value"`
	Type      string    `json:"type"`
	ID1       string    `json:"id1"`
	ID2       int       `json:"id2"`
	Timestamp time.Time `json:"timestamp"`
}

var (
	mu           sync.Mutex
	generateFreq time.Duration = time.Second
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Endpoint to change the frequency of data generation
	e.POST("/frequency", func(c echo.Context) error {
		freq := struct {
			Frequency int `json:"frequency"`
		}{}
		if err := c.Bind(&freq); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		mu.Lock()
		generateFreq = time.Duration(freq.Frequency) * time.Millisecond
		mu.Unlock()

		return c.NoContent(http.StatusOK)
	})

	// Start data generation
	go generateData()

	e.Logger.Fatal(e.Start(":8080"))
}

func generateData() {
	for {
		mu.Lock()
		time.Sleep(generateFreq)
		mu.Unlock()

		data := SensorData{
			Value:     rand.Float64() * 100,
			Type:      "Temperature",
			ID1:       string(randomCapitalChar()),
			ID2:       rand.Intn(100),
			Timestamp: time.Now(),
		}
		fmt.Println(fmt.Printf("%+v", data))

		// Todo: send data to other services
	}
}
func randomCapitalChar() rune {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(26)
	return rune('A' + randomIndex)
}
