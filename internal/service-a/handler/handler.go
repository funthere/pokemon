package handler

import (
	"net/http"
	"time"

	"github.com/funthere/pokemon/internal/service-a/service"
	"github.com/labstack/echo/v4"
)

type FrequencyRequest struct {
	Frequency int `json:"frequency"`
}

func SetFrequency(c echo.Context) error {
	req := new(FrequencyRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	service.UpdateFrequency(req.Frequency)
	return c.JSON(http.StatusOK, map[string]string{"status": "frequency updated"})
}

func GenerateData() {
	for {
		service.GenerateSensorData()
		time.Sleep(time.Millisecond * time.Duration(service.GetFrequency()))
	}
}
