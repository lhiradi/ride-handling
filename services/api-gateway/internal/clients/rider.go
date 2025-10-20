package clients

import (
	riderv1 "github.com/lhiradi/ride-handling/proto/rider/v1"
	"google.golang.org/grpc"
)

func NewRiderClient(conn *grpc.ClientConn) riderv1.RiderServiceClient {
	return riderv1.NewRiderServiceClient(conn)
}
