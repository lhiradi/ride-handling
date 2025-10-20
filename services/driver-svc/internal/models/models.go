package models

import "gorm.io/gorm"

type Driver struct {
	ID      string `gorm:"primarykey"`
	Name    string
	Phone   string
	Vechile string
	Status  string // online or offline
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Driver{})
}
