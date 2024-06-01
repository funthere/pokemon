package main

import (
	"github.com/funthere/pokemon/internal/service-a/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	go handler.GenerateData()

	e := echo.New()
	e.POST("/set-frequency", handler.SetFrequency)

	e.Logger.Fatal(e.Start(":8081"))
}
