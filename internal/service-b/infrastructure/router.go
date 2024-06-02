package infrastructure

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/funthere/pokemon/internal/service-b/docs" // import generated docs
	"github.com/funthere/pokemon/internal/service-b/handler"
	mw "github.com/funthere/pokemon/internal/service-b/middleware"
	"github.com/funthere/pokemon/internal/service-b/usecase"
)

func NewRouter(sensorUsecase usecase.SensorUsecase) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Basic Auth Middleware
	e.Use(mw.BasicAuth("admin", "password"))

	sensorHandler := handler.NewSensorHandler(sensorUsecase)

	e.GET("/data", sensorHandler.Fetch)
	e.DELETE("/data", sensorHandler.Delete)
	e.PUT("/data", sensorHandler.Update)

	// Swagger endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return e
}
