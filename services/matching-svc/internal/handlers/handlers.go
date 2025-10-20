package handlers

import (
	"context"
	"log"

	"github.com/google/uuid"
	matchingv1 "github.com/lhiradi/ride-handling/proto/matching/v1"
	"github.com/lhiradi/ride-handling/services/matching-svc/internal/client"
	"github.com/redis/go-redis/v9"
)

type MatchingServer struct {
	matchingv1.UnimplementedMatchingServiceServer
	Rdb        *redis.Client
	TripClient client.TripClient
}

func (m *MatchingServer) Match(ctx context.Context, req *matchingv1.MatchRequest) (*matchingv1.MatchResponse, error) {
	res, err := m.Rdb.GeoRadius(ctx, "active_drivers", req.Pickup.Lon, req.Pickup.Lat,
		&redis.GeoRadiusQuery{
			Radius: float64(req.RadiusM),
			Unit:   "m",
			Count:  int(req.Limit),
		}).Result()
	if err != nil {
		return nil, err
	}

	invites := []*matchingv1.Invitation{}
	for _, d := range res {
		invites = append(invites, &matchingv1.Invitation{
			Id:               uuid.NewString(),
			TripId:           req.TripId,
			DriverId:         d.Name,
			Status:           "pending",
			ExpiresAtUnixSec: 0, // optional
		})
	}
	return &matchingv1.MatchResponse{Invites: invites}, nil
}

func (m *MatchingServer) AcceptInvitation(ctx context.Context, req *matchingv1.AcceptInvitationRequest) (*matchingv1.AcceptInvitationResponse, error) {
	err := m.TripClient.AssignDriver(ctx, req.TripId, req.DriverId)
	if err != nil {
		log.Printf("failed to assign driver: %v", err)
		return &matchingv1.AcceptInvitationResponse{Accepted: false}, err
	}
	return &matchingv1.AcceptInvitationResponse{Accepted: true}, nil
}
