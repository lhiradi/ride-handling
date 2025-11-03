package models

type Driver struct {
	ID           string `gorm:"primarykey"`
	Name         string
	Phone        string
	Vechile      string
	Status       string // online or offline
	PasswordHash string
}
