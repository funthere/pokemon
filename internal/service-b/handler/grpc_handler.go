package handler

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/funthere/pokemon/internal/service-b/service"
	pb "github.com/funthere/pokemon/proto"
)

type grpcHandler struct {
	service *service.Service
	pb.UnimplementedSensorServiceServer
}

func NewGrpcHandler(db *sql.DB) *grpcHandler {
	return &grpcHandler{
		service: service.NewService(db),
	}
}

// Used in gRPC
func (h *grpcHandler) SendSensorData(ctx context.Context, req *pb.SensorData) (*pb.SensorResponse, error) {
	fmt.Printf("%+v\n", req)
	err := h.service.SaveData(float64(req.Value), req.Type, req.Id1, int(req.Id2), req.Timestamp)
	if err != nil {
		return &pb.SensorResponse{Status: "error"}, err
	}
	return &pb.SensorResponse{Status: "data received"}, nil
}
