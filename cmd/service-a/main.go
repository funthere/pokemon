package main

import (
	"log"
	"net/http"

	"github.com/funthere/pokemon/internal/service-a/handler"
	"github.com/funthere/pokemon/internal/service-a/service"
	pb "github.com/funthere/pokemon/proto"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// gRPC client
	conn, err := grpc.NewClient("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewSensorServiceClient(conn)

	// Generate sensor's data and send them to client
	go service.GenerateData(client)

	// Handle REST API
	e := echo.New()
	e.POST("/set-frequency", handler.SetFrequency)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8081"))
}
