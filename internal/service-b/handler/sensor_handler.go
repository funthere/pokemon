package handler

import (
	"net/http"
	"strconv"

	"github.com/funthere/pokemon/internal/service-b/domain"
	"github.com/funthere/pokemon/internal/service-b/usecase"

	"github.com/labstack/echo/v4"
)

type SensorHandler struct {
	SensorUsecase usecase.SensorUsecase
}

func NewSensorHandler(e *echo.Echo, us usecase.SensorUsecase) {
	handler := &SensorHandler{
		SensorUsecase: us,
	}

	e.GET("/data", handler.Fetch)
	e.DELETE("/data", handler.Delete)
	e.PUT("/data", handler.Update)
}

func (h *SensorHandler) Fetch(c echo.Context) error {
	id1 := c.QueryParam("id1")
	id2 := c.QueryParam("id2")
	start := c.QueryParam("start")
	end := c.QueryParam("end")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("size"))

	data, err := h.SensorUsecase.Fetch(id1, id2, start, end, page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (h *SensorHandler) Delete(c echo.Context) error {
	id1 := c.QueryParam("id1")
	id2 := c.QueryParam("id2")
	start := c.QueryParam("start")
	end := c.QueryParam("end")

	rowsAffected, err := h.SensorUsecase.Delete(id1, id2, start, end)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]int64{"deleted": rowsAffected})
}

func (h *SensorHandler) Update(c echo.Context) error {
	id1 := c.QueryParam("id1")
	id2 := c.QueryParam("id2")
	start := c.QueryParam("start")
	end := c.QueryParam("end")

	var data domain.SensorData
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	rowsAffected, err := h.SensorUsecase.Update(data, id1, id2, start, end)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]int64{"updated": rowsAffected})
}
