package db

import (
	"log"
	"sync"

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
