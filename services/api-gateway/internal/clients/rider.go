package clients

import (
	"context"

	riderv1 "github.com/lhiradi/ride-handling/proto/rider/v1"
	"google.golang.org/grpc"
)

type RiderClient struct {
	api riderv1.RiderServiceClient
}

func NewRiderClient(conn *grpc.ClientConn) *RiderClient {
	return &RiderClient{api: riderv1.NewRiderServiceClient(conn)}
}

func (c *RiderClient) CreateRider(ctx context.Context, req *riderv1.CreateRiderRequest) (*riderv1.CreateRiderResponse, error) {
	return c.api.CreateRider(ctx, req)
}
func (c *RiderClient) GetRider(ctx context.Context, req *riderv1.GetRiderRequest) (*riderv1.GetRiderResponse, error) {
	return c.api.GetRider(ctx, req)
}
