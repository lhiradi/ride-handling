package clients

import (
	matchingv1 "github.com/lhiradi/ride-handling/proto/matching/v1"
	"google.golang.org/grpc"
)

func NewMatchingClient(conn *grpc.ClientConn) matchingv1.MatchingServiceClient {
	return matchingv1.NewMatchingServiceClient(conn)
}
