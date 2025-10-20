package clients

import (
	driverv1 "github.com/lhiradi/ride-handling/proto/driver/v1"
	"google.golang.org/grpc"
)

func NewDriverClient(conn *grpc.ClientConn) driverv1.DriverServiceClient {
	return driverv1.NewDriverServiceClient(conn)
}
