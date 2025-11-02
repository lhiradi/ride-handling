package services

import (
	"context"
	"strings"

	riderv1 "github.com/lhiradi/ride-handling/proto/rider/v1"
	"github.com/lhiradi/ride-handling/services/api-gateway/internal/clients"
)

type RiderService struct {
	Client *clients.RiderClient
}

func (s *RiderService) CreateRider(ctx context.Context, name, phone, email, language string) (*riderv1.CreateRiderResponse, error) {
	name = strings.TrimSpace(name)
	phone = strings.TrimSpace(phone)
	email = strings.TrimSpace(email)
	language = strings.TrimSpace(language)

	req := &riderv1.CreateRiderRequest{
		Name:     name,
		Phone:    phone,
		Email:    email,
		Language: language,
	}
	return s.Client.CreateRider(ctx, req)
}

func (s *RiderService) GetRider(ctx context.Context, riderID string) (*riderv1.GetRiderResponse, error) {
	req := &riderv1.GetRiderRequest{RiderId: strings.TrimSpace(riderID)}
	return s.Client.GetRider(ctx, req)
}
