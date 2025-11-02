// services/api-gateway/cmd/main.go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lhiradi/ride-handling/services/api-gateway/internal/clients"
	"github.com/lhiradi/ride-handling/services/api-gateway/internal/router"
	"github.com/lhiradi/ride-handling/services/api-gateway/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	app := fiber.New()

	// gRPC connections
	driverConn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to DriverService: %v", err)
	}
	riderConn, err := grpc.Dial("localhost:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to RiderService: %v", err)
	}
	tripConn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to TripService: %v", err)
	}
	matchingConn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to MatchingService: %v", err)
	}

	// Clients
	driverClient := clients.NewDriverClient(driverConn)
	riderClient := clients.NewRiderClient(riderConn)
	tripClient := clients.NewTripClient(tripConn)
	matchingClient := clients.NewMatchingClient(matchingConn)

	// Services
	driverSvc := &services.DriverService{Client: driverClient}
	riderSvc := &services.RiderService{Client: riderClient}
	tripSvc := &services.TripService{Client: tripClient}
	matchingSvc := &services.MatchingService{Client: matchingClient}

	// Router setup
	router.Setup(app, driverSvc, riderSvc, tripSvc, matchingSvc)

	// Start server
	log.Println("API Gateway listening on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("failed to start API Gateway: %v", err)
	}
}
