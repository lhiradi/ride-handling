package clients

import (
	"context"

	matchingv1 "github.com/lhiradi/ride-handling/proto/matching/v1"
	"google.golang.org/grpc"
)

type MatchingClient struct {
	api matchingv1.MatchingServiceClient
}

func NewMatchingClient(conn *grpc.ClientConn) *MatchingClient {
	return &MatchingClient{api: matchingv1.NewMatchingServiceClient(conn)}
}

func (c *MatchingClient) Match(ctx context.Context, req *matchingv1.MatchRequest) (*matchingv1.MatchResponse, error) {
	return c.api.Match(ctx, req)
}
func (c *MatchingClient) AcceptInvitation(ctx context.Context, req *matchingv1.AcceptInvitationRequest) (*matchingv1.AcceptInvitationResponse, error) {
	return c.api.AcceptInvitation(ctx, req)
}
