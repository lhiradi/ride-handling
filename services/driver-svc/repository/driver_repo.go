package repository

import (
	"context"
	"errors"

	"github.com/lhiradi/ride-handling/services/driver-svc/internal/models"
	"gorm.io/gorm"
)

type DriverRepository interface {
	Create(ctx context.Context, driver *models.Driver) error
	GetByID(ctx context.Context, id string) (*models.Driver, error)
	UpdateStatus(ctx context.Context, id string, status string) error
	ListOnline(ctx context.Context) ([]*models.Driver, error)
}

type driverRepo struct {
	DB *gorm.DB
}

func NewDriverRepo(db *gorm.DB) DriverRepository {
	return &driverRepo{DB: db}
}

func (r *driverRepo) Create(ctx context.Context, driver *models.Driver) error {
	return r.DB.WithContext(ctx).Create(driver).Error
}

func (r *driverRepo) GetByID(ctx context.Context, id string) (*models.Driver, error) {
	var driver models.Driver
	if err := r.DB.WithContext(ctx).First(&driver, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &driver, nil
}

func (r *driverRepo) UpdateStatus(ctx context.Context, id string, status string) error {
	return r.DB.WithContext(ctx).
		Model(&models.Driver{}).
		Where("id = ?", id).
		Update("status", status).Error
}

func (r *driverRepo) ListOnline(ctx context.Context) ([]*models.Driver, error) {
	var drivers []*models.Driver
	if err := r.DB.WithContext(ctx).
		Where("status = ?", "online").
		Find(&drivers).Error; err != nil {
		return nil, err
	}
	return drivers, nil
}
