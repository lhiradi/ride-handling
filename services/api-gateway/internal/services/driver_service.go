package services

import (
	"context"

	commonv1 "github.com/lhiradi/ride-handling/proto/common/v1"
	driverv1 "github.com/lhiradi/ride-handling/proto/driver/v1"
	"github.com/lhiradi/ride-handling/services/api-gateway/internal/clients"
)

type DriverService struct {
	Client *clients.DriverClient
}

func (s *DriverService) SetStatus(ctx context.Context, driverID, status string) error {
	req := &driverv1.SetStatusRequest{
		DriverId: driverID,
		Status:   status,
	}
	return s.Client.SetStatus(ctx, req)
}

func (s *DriverService) Heartbeat(ctx context.Context, driverID string, lat, lon float64, ts int64) error {
	req := &driverv1.HeartbeatRequest{
		DriverId:  driverID,
		Location:  &commonv1.GeoPoint{Lat: lat, Lon: lon},
		AtUnixSec: ts,
	}
	return s.Client.Heartbeat(ctx, req)
}
