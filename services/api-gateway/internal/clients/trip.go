package clients

import (
	"context"

	tripv1 "github.com/lhiradi/ride-handling/proto/trip/v1"
	"google.golang.org/grpc"
)

type TripClient struct {
	api tripv1.TripServiceClient
}

func NewTripClient(conn *grpc.ClientConn) *TripClient {
	return &TripClient{api: tripv1.NewTripServiceClient(conn)}
}

func (c *TripClient) CreateTrip(ctx context.Context, req *tripv1.CreateTripRequest) (*tripv1.CreateTripResponse, error) {
	return c.api.CreateTrip(ctx, req)
}
