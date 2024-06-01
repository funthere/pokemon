package handler

import (
	"database/sql"
	"net/http"

	"github.com/funthere/pokemon/helper"
	"github.com/funthere/pokemon/internal/service-b/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *service.Service
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		service: service.NewService(db),
	}
}

type SensorDataRequest struct {
	Value     float64 `json:"value"`
	Type      string  `json:"type"`
	ID1       string  `json:"id1"`
	ID2       int     `json:"id2"`
	Timestamp string  `json:"timestamp"`
}

func (h *Handler) GetData(c echo.Context) error {
	page := helper.StringToUint32(c.QueryParam("page"))
	size := helper.StringToUint32(c.QueryParam("size"))

	data, err := h.service.Fetch(page, size)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, data)
}
