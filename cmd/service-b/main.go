package main

import (
	"log"
	"net"

	"github.com/funthere/pokemon/internal/service-b/handler"
	"github.com/funthere/pokemon/internal/service-b/infrastructure"
	"github.com/funthere/pokemon/internal/service-b/repository"
	"github.com/funthere/pokemon/internal/service-b/usecase"
	pb "github.com/funthere/pokemon/proto"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
)

func main() {
	// REST API
	db := infrastructure.NewMysqlDB()
	defer db.Close()

	sensorRepo := repository.NewMysqlSensorRepository(db)
	sensorUsecase := usecase.NewSensorUsecase(sensorRepo)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	handler.NewSensorHandler(e, sensorUsecase)

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
