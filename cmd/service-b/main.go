package main

import (
	"github.com/funthere/pokemon/internal/service-b/handler"
	"github.com/funthere/pokemon/pkg/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := database.Connect()
	defer db.Close()

	h := handler.NewHandler(db)

	e.GET("/data", h.GetData)

	e.Logger.Fatal(e.Start(":8082"))
}
