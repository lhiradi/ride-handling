package clients

import (
	"context"

	driverv1 "github.com/lhiradi/ride-handling/proto/driver/v1"
	"google.golang.org/grpc"
)

type DriverClient struct {
	api driverv1.DriverServiceClient
}

func NewDriverClient(conn *grpc.ClientConn) *DriverClient {
	return &DriverClient{api: driverv1.NewDriverServiceClient(conn)}
}

func (c *DriverClient) SetStatus(ctx context.Context, req *driverv1.SetStatusRequest) error {
	_, err := c.api.SetStatus(ctx, req)
	return err
}

func (c *DriverClient) Heartbeat(ctx context.Context, req *driverv1.HeartbeatRequest) error {
	_, err := c.api.Heartbeat(ctx, req)
	return err
}
