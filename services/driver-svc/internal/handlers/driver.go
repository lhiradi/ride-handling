package handlers

import (
	"context"

	driverv1 "github.com/lhiradi/ride-handling/proto/driver/v1"
	"github.com/lhiradi/ride-handling/services/driver-svc/internal/models"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type DriverServer struct {
	driverv1.UnimplementedDriverServiceServer
	DB *gorm.DB
}

func (s *DriverServer) SetStatus(ctx context.Context, req *driverv1.SetStatusRequest) (*emptypb.Empty, error) {
	if err := s.DB.Model(&models.Driver{}).Where("id = ?", req.DriverId).Update("status", req.Status).Error; err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *DriverServer) Heartbeat(ctx context.Context, req *driverv1.HeartbeatRequest) (*emptypb.Empty, error) {
	// In a real system, push location into Redis GEO
	// For now, just log or no-op
	return &emptypb.Empty{}, nil
}
