package clients

import (
	tripv1 "github.com/lhiradi/ride-handling/proto/trip/v1"
	"google.golang.org/grpc"
)

func NewTripClient(conn *grpc.ClientConn) tripv1.TripServiceClient {
	return tripv1.NewTripServiceClient(conn)
}
