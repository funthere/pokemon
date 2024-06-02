package main

import (
	"log"
	"net"

	_ "github.com/funthere/pokemon/internal/service-b/docs" // import generated docs
	"github.com/funthere/pokemon/internal/service-b/handler"
	"github.com/funthere/pokemon/internal/service-b/infrastructure"
	"github.com/funthere/pokemon/internal/service-b/repository"
	"github.com/funthere/pokemon/internal/service-b/usecase"
	pb "github.com/funthere/pokemon/proto"
	"google.golang.org/grpc"
)

// @title Microservice B API
// @version 1.0
// @description This is the API documentation for Microservice B
// @host localhost:8082
// @BasePath /
func main() {
	// REST API
	db := infrastructure.NewMysqlDB()
	defer db.Close()

	sensorRepo := repository.NewMysqlSensorRepository(db)
	sensorUsecase := usecase.NewSensorUsecase(sensorRepo)
	e := infrastructure.NewRouter(sensorUsecase)

	go func() {
		e.Logger.Fatal(e.Start(":8082"))
	}()

	// gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	h := handler.NewGrpcHandler(db)
	pb.RegisterSensorServiceServer(grpcServer, h)

	log.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
