package main

import (
	"log"
	"net"

	riderv1 "github.com/lhiradi/ride-handling/proto/rider/v1"
	"github.com/lhiradi/ride-handling/services/rider-svc/db"
	"github.com/lhiradi/ride-handling/services/rider-svc/internal/handlers"
	"github.com/lhiradi/ride-handling/services/rider-svc/internal/models"
	"github.com/lhiradi/ride-handling/services/rider-svc/internal/utils"
	"google.golang.org/grpc"
)

func main() {
	dsn := utils.GetDsn()
	db.InitDB(dsn)
	database := db.Get()

	if err := models.AutoMigrate(database); err != nil {
		log.Fatal(err)
	}

	lis, _ := net.Listen("tcp", ":50054")
	s := grpc.NewServer()
	riderv1.RegisterRiderServiceServer(s, &handlers.RiderServer{DB: database})
	log.Println("Rider service listening on :50054")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
