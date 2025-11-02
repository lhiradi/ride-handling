package services

import (
	"context"

	commonv1 "github.com/lhiradi/ride-handling/proto/common/v1"
	tripv1 "github.com/lhiradi/ride-handling/proto/trip/v1"
	"github.com/lhiradi/ride-handling/services/api-gateway/internal/clients"
)

type TripService struct {
	Client *clients.TripClient
}

func (s *TripService) CreateTrip(ctx context.Context, riderID string, pickupLat, pickupLon, dropLat, dropLon float64) (*tripv1.CreateTripResponse, error) {
	req := &tripv1.CreateTripRequest{
		RiderId: riderID,
		Pickup:  &commonv1.GeoPoint{Lat: pickupLat, Lon: pickupLon},
		Dropoff: &commonv1.GeoPoint{Lat: dropLat, Lon: dropLon},
	}
	return s.Client.CreateTrip(ctx, req)
}
