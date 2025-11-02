package client

import (
	"context"
	"log"

	driverv1 "github.com/lhiradi/ride-handling/proto/driver/v1"
	"google.golang.org/grpc"
)

type DriverClient struct {
	api driverv1.DriverServiceClient
}

func NewDriverServiceClient(conn *grpc.ClientConn) *DriverClient {
	return &DriverClient{api: driverv1.NewDriverServiceClient(conn)}
}

func (c *DriverClient) FindNearbyDrivers(ctx context.Context, lat, lon, radius float64) ([]string, error) {
	resp, err := c.api.FindNearbyDrivers(ctx, &driverv1.FindNearbyDriversRequest{Lat: lat, Lon: lon, RadiusKm: radius})
	if err != nil {
		return nil, err
	}
	log.Printf("Found %d drivers", len(resp.DriverIds))
	return resp.DriverIds, nil
}
