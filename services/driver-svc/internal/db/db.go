package db

import (
	"log"
	"sync"

	"github.com/lhiradi/ride-handling/services/driver-svc/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	once     sync.Once
)

func InitDB(dsn string) {
	once.Do(func() {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}
		if err := db.AutoMigrate(&models.Driver{}); err != nil {
			log.Fatalf("failed to migrate: %v", err)
		}
		instance = db
	})
}

// Get returns the singleton DB instance
func Get() *gorm.DB {
	if instance == nil {
		log.Fatal("database not initialized, call db.Init(dsn) first")
	}
	return instance
}
