package main

import (
	"log"
	"net"

	driverv1 "github.com/lhiradi/ride-handling/proto/driver/v1"
	"github.com/lhiradi/ride-handling/services/driver-svc/internal/db"
	"github.com/lhiradi/ride-handling/services/driver-svc/internal/handlers"
	"github.com/lhiradi/ride-handling/services/driver-svc/repository"

	"github.com/lhiradi/ride-handling/services/driver-svc/internal/utils"
	"google.golang.org/grpc"
)

func main() {
	dsn := utils.GetDsn()

	// Initialize DB
	db.InitDB(dsn)
	database := db.Get()

	driverRepo := repository.NewDriverRepo(database)
	grpcServer := grpc.NewServer()

	driverv1.RegisterDriverServiceServer(grpcServer, &handlers.DriverServer{
		Repo: driverRepo,
	})

	// Start listener
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen on port 50052: %v", err)
	}

	log.Println("Driver service listening on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}
