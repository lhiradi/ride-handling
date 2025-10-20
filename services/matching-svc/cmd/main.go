package main

import (
	"log"
	"net"

	matchingv1 "github.com/lhiradi/ride-handling/proto/matching/v1"
	"github.com/lhiradi/ride-handling/services/matching-svc/internal/client"
	"github.com/lhiradi/ride-handling/services/matching-svc/internal/handlers"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	tripConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to connect to TripService:", err)
	}
	tripClient := client.NewTripClient(tripConn)

	lis, _ := net.Listen("tcp", ":50053")
	s := grpc.NewServer()
	matchingv1.RegisterMatchingServiceServer(s, &handlers.MatchingServer{
		Rdb:        rdb,
		TripClient: tripClient,
	})
	log.Println("Matching service listening on :50053")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
