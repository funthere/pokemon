package handler

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/funthere/pokemon/internal/service-b/service"
	pb "github.com/funthere/pokemon/proto"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *service.Service
	pb.UnimplementedSensorServiceServer
}

func NewGrpcHandler(db *sql.DB) *Handler {
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

// Basic auth validator
func BasicAuthValidator(username, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "password" {
		return true, nil
	}
	return false, nil
}
