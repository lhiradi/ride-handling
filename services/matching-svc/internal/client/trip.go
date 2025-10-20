package client

import (
	"context"

	commonv1 "github.com/lhiradi/ride-handling/proto/common/v1"
	tripv1 "github.com/lhiradi/ride-handling/proto/trip/v1"
	"google.golang.org/grpc"
)

type TripClient struct {
	tripv1.TripServiceClient
}

func NewTripClient(conn *grpc.ClientConn) TripClient {
	return TripClient{TripServiceClient: tripv1.NewTripServiceClient(conn)}
}

func (c TripClient) AssignDriver(ctx context.Context, tripID, driverID string) error {
	_, err := c.TripServiceClient.UpdateTripStatus(ctx, &tripv1.UpdateTripStatusRequest{
		TripId:   tripID,
		DriverId: driverID,
		Status:   commonv1.TripStatus_ASSIGNED,
	})
	return err
}
