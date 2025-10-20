package handlers

import (
	"context"

	"github.com/google/uuid"
	riderv1 "github.com/lhiradi/ride-handling/proto/rider/v1"
	"github.com/lhiradi/ride-handling/services/rider-svc/internal/models"
	"gorm.io/gorm"
)

type RiderServer struct {
	riderv1.UnimplementedRiderServiceServer
	DB *gorm.DB
}

func (s *RiderServer) CreateRider(ctx context.Context, req *riderv1.CreateRiderRequest) (*riderv1.CreateRiderResponse, error) {
	id := uuid.NewString()
	rider := &models.Rider{
		ID: id, Name: req.Name, Phone: req.Phone, Email: req.Email, Language: req.Language,
	}
	if err := s.DB.Create(rider).Error; err != nil {
		return nil, err
	}
	return &riderv1.CreateRiderResponse{RiderId: id}, nil
}

func (s *RiderServer) GetRider(ctx context.Context, req *riderv1.GetRiderRequest) (*riderv1.GetRiderResponse, error) {
	var rider models.Rider
	if err := s.DB.First(&rider, "id = ?", req.RiderId).Error; err != nil {
		return nil, err
	}
	return &riderv1.GetRiderResponse{
		Name: rider.Name, Phone: rider.Phone, Email: rider.Email, Language: rider.Language,
	}, nil
}
