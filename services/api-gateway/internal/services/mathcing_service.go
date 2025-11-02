package services

import (
	"context"

	commonv1 "github.com/lhiradi/ride-handling/proto/common/v1"
	matchingv1 "github.com/lhiradi/ride-handling/proto/matching/v1"
	"github.com/lhiradi/ride-handling/services/api-gateway/internal/clients"
)

type MatchingService struct {
	Client *clients.MatchingClient
}

func (s *MatchingService) MatchTrip(ctx context.Context, tripID string, lat, lon, radius float64, limit int32) (*matchingv1.MatchResponse, error) {
	return s.Client.Match(ctx, &matchingv1.MatchRequest{
		TripId:  tripID,
		Pickup:  &commonv1.GeoPoint{Lat: lat, Lon: lon},
		RadiusM: int32(radius * 1000),
		Limit:   limit,
	})
}
func (s *MatchingService) AcceptInvitation(ctx context.Context, invitationID, tripID, driverID string) (*matchingv1.AcceptInvitationResponse, error) {
	return s.Client.AcceptInvitation(ctx, &matchingv1.AcceptInvitationRequest{
		InvitationId: invitationID,
		TripId:       tripID,
		DriverId:     driverID,
	})
}
