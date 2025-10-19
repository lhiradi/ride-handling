package handlers

import (
	"context"

	"github.com/google/uuid"
	commonv1 "github.com/lhiradi/ride-handling/proto/common/v1"
	tripv1 "github.com/lhiradi/ride-handling/proto/trip/v1"
	"github.com/lhiradi/ride-handling/services/trip-svc/internal/models"
	"gorm.io/gorm"
)

type TripServer struct {
	tripv1.UnimplementedTripServiceServer
	DB *gorm.DB
}

func (s *TripServer) CreateTrip(ctx context.Context, req *tripv1.CreateTripRequest) (*tripv1.CreateTripResponse, error) {
	id := uuid.NewString()
	trip := &models.Trip{ID: id,
		RiderID:   req.RiderId,
		PickupLat: req.Pickup.Lat,
		PickupLon: req.Pickup.Lon,
		DropLat:   req.Dropoff.Lat,
		DropLon:   req.Dropoff.Lon,
		Status:    int32(commonv1.TripStatus_REQUESTED)}

	if err := s.DB.Create(trip).Error; err != nil {
		return nil, err
	}

	return &tripv1.CreateTripResponse{
		Trip: &tripv1.Trip{
			Id: id, RiderId: req.RiderId,
			Pickup:  &commonv1.GeoPoint{Lat: req.Pickup.Lat, Lon: req.Pickup.Lon},
			Dropoff: &commonv1.GeoPoint{Lat: req.Dropoff.Lat, Lon: req.Dropoff.Lon},
			Status:  commonv1.TripStatus_REQUESTED,
		},
	}, nil
}
