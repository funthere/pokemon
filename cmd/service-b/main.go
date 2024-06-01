package main

import (
	"log"
	"net"

	"github.com/funthere/pokemon/internal/service-b/handler"
	"github.com/funthere/pokemon/pkg/database"
	pb "github.com/funthere/pokemon/proto"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
)

func main() {
	// REST API
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := database.Connect()
	defer db.Close()

	h := handler.NewHandler(db)

	e.GET("/data", h.GetData)

	go func() {
		e.Logger.Fatal(e.Start(":8082"))
	}()

	// gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSensorServiceServer(grpcServer, h)

	log.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
