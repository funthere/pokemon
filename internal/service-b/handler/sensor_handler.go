package handler

import (
	"net/http"
	"strconv"

	"github.com/funthere/pokemon/internal/service-b/domain"
	"github.com/funthere/pokemon/internal/service-b/usecase"

	_ "github.com/funthere/pokemon/internal/service-b/docs" // import generated docs

	"github.com/labstack/echo/v4"
)

type SensorHandler interface {
	Fetch(c echo.Context) error
	Delete(c echo.Context) error
	Update(c echo.Context) error
}

type sensorHandler struct {
	SensorUsecase usecase.SensorUsecase
}

// NewSensorHandler creates a new sensor handler
func NewSensorHandler(sensorUsecase usecase.SensorUsecase) SensorHandler {
	return &sensorHandler{
		SensorUsecase: sensorUsecase,
	}
}

// Fetch retrieves sensor data
// @Summary Fetch sensor data
// @Description Get sensor data by ID and timestamp
// @Tags Sensor
// @Accept  json
// @Produce  json
// @Param id1 query string false "ID1"
// @Param id2 query string false "ID2"
// @Param start query string false "Start Time"
// @Param end query string false "End Time"
// @Param page query int false "Page number"
// @Param size query int false "Page size"
// @Success 200 {array} domain.SensorData
// @Router /data [get]
func (h *sensorHandler) Fetch(c echo.Context) error {
	id1 := c.QueryParam("id1")
	id2 := c.QueryParam("id2")
	start := c.QueryParam("start")
	end := c.QueryParam("end")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	size, _ := strconv.Atoi(c.QueryParam("size"))
	pagination := domain.Pagination{
		Page: uint(page),
		Size: uint(size),
	}

	data, err := h.SensorUsecase.Fetch(id1, id2, start, end, &pagination)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	res := map[string]any{
		"pagination": pagination,
		"data":       data,
	}

	return c.JSON(http.StatusOK, res)
}

// Delete removes sensor data
// @Summary Delete sensor data
// @Description Delete sensor data by ID and timestamp
// @Tags Sensor
// @Accept  json
// @Produce  json
// @Param id1 query string false "ID1"
// @Param id2 query string false "ID2"
// @Param start query string false "Start Time"
// @Param end query string false "End Time"
// @Success 200 {object} map[string]int64
// @Router /data [delete]
func (h *sensorHandler) Delete(c echo.Context) error {
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

// Update modifies sensor data
// @Summary Update sensor data
// @Description Update sensor data by ID and timestamp
// @Tags Sensor
// @Accept  json
// @Produce  json
// @Param id1 query string false "ID1"
// @Param id2 query string false "ID2"
// @Param start query string false "Start Time"
// @Param end query string false "End Time"
// @Param data body domain.SensorDataUpdateReq true "Sensor Data"
// @Success 200 {object} map[string]int64
// @Router /data [put]
func (h *sensorHandler) Update(c echo.Context) error {
	id1 := c.QueryParam("id1")
	id2 := c.QueryParam("id2")
	start := c.QueryParam("start")
	end := c.QueryParam("end")

	var data domain.SensorDataUpdateReq
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	rowsAffected, err := h.SensorUsecase.Update(data, id1, id2, start, end)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]int64{"updated": rowsAffected})
}
