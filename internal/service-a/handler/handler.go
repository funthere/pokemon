package handler

import (
	"fmt"
	"net/http"

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
	fmt.Println("======Freq updated=====")
	return c.JSON(http.StatusOK, map[string]string{"status": "frequency updated"})
}
