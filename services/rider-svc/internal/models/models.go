package models

import "gorm.io/gorm"

type Rider struct {
	ID       string `gorm:"primaryKey"`
	Name     string
	Phone    string
	Email    string
	Language string
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Rider{})
}
