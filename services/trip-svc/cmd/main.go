package main

import (
	"log"
	"net"

	tripv1 "github.com/lhiradi/ride-handling/proto/trip/v1"
	"github.com/lhiradi/ride-handling/services/trip-svc/internal/db"
	"github.com/lhiradi/ride-handling/services/trip-svc/internal/handlers"
	"github.com/lhiradi/ride-handling/services/trip-svc/internal/models"
	"github.com/lhiradi/ride-handling/services/trip-svc/utils"
	"google.golang.org/grpc"
)

func main() {
	dsn := utils.GetDsn()

	db.InitDB(dsn)
	database := db.Get()

	if err := models.AutoMigrate(database); err != nil {
		log.Fatal(err)
	}

	// Start gRPC server
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	tripv1.RegisterTripServiceServer(s, &handlers.TripServer{DB: database})
	log.Println("Trip service listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
