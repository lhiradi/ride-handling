package handlers

import (
	"context"
	"log"

	driverv1 "github.com/lhiradi/ride-handling/proto/driver/v1"
	"github.com/lhiradi/ride-handling/services/driver-svc/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DriverServer struct {
	driverv1.UnimplementedDriverServiceServer
	Repo repository.DriverRepository
}

func (s *DriverServer) SetStatus(ctx context.Context, req *driverv1.SetStatusRequest) (*emptypb.Empty, error) {
	if req.DriverId == "" {
		return nil, status.Error(codes.InvalidArgument, ("driver_id is required"))
	}
	if req.Status == "" {
		return nil, status.Error(codes.InvalidArgument, ("status is required"))
	}

	if err := s.Repo.UpdateStatus(ctx, req.DriverId, req.Status); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// Heartbeat records driver’s location + timestamp
func (s *DriverServer) Heartbeat(ctx context.Context, req *driverv1.HeartbeatRequest) (*emptypb.Empty, error) {
	if req.DriverId == "" {
		return nil, status.Error(codes.InvalidArgument, "driver_id is required")
	}
	if req.Location == nil {
		return nil, status.Error(codes.InvalidArgument, "location is required")
	}

	// Later: publish to Redis GEO for fast matching
	log.Printf("Heartbeat from driver %s at (%.6f, %.6f) ts=%d",
		req.DriverId, req.Location.Lat, req.Location.Lon, req.AtUnixSec)

	return &emptypb.Empty{}, nil
}
