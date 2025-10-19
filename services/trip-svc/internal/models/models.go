package models

import "gorm.io/gorm"

type Trip struct {
	ID        string `gorm:"primarykey"`
	RiderID   string
	DriverID  string
	PickupLat float64
	PickupLon float64
	DropLat   float64
	DropLon   float64
	Status    int32
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Trip{})
}
