package handler

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/funthere/pokemon/helper"
	"github.com/funthere/pokemon/internal/service-b/service"
	pb "github.com/funthere/pokemon/proto"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *service.Service
	pb.UnimplementedSensorServiceServer
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

// Used in gRPC
func (h *Handler) SendSensorData(ctx context.Context, req *pb.SensorData) (*pb.SensorResponse, error) {
	fmt.Printf("%+v\n", req)
	err := h.service.SaveData(float64(req.Value), req.Type, req.Id1, int(req.Id2), req.Timestamp)
	if err != nil {
		return &pb.SensorResponse{Status: "error"}, err
	}
	return &pb.SensorResponse{Status: "data received"}, nil
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
