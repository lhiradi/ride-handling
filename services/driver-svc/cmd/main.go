package main

import (
	"log"
	"net"

	driverv1 "github.com/lhiradi/ride-handling/proto/driver/v1"
	"github.com/lhiradi/ride-handling/services/driver-svc/internal/db"
	"github.com/lhiradi/ride-handling/services/driver-svc/internal/handlers"
	"github.com/lhiradi/ride-handling/services/driver-svc/internal/models"
	"github.com/lhiradi/ride-handling/services/driver-svc/internal/utils"
	"google.golang.org/grpc"
)

func main() {
	dsn := utils.GetDsn()
	db.InitDB(dsn)
	database := db.Get()

	if err := models.AutoMigrate(database); err != nil {
		log.Fatal(err)
	}

	lis, _ := net.Listen("tcp", ":50052")
	s := grpc.NewServer()
	driverv1.RegisterDriverServiceServer(s, &handlers.DriverServer{DB: database})
	log.Println("Driver service listening on :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
