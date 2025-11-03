package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
	Port                string
	JWTSecret           string
	JWTTTL              time.Duration
	RedisAddr           string
	PostgresDSN         string
	TripServiceAddr     string
	DriverServiceAddr   string
	RiderServiceAddr    string
	MatchingServiceAddr string
	Env                 string
}

func Load() *Config {
	cfg := &Config{
		Port:                getEnv("PORT", "8080"),
		JWTSecret:           getEnv("JWT_SECRET", "change-me"),
		JWTTTL:              mustDuration(getEnv("JWT_TTL", "24h")),
		RedisAddr:           getEnv("REDIS_ADDR", "localhost:6379"),
		PostgresDSN:         getEnv("POSTGRES_DSN", "postgres://postgres:postgres@localhost:5432/trip?sslmode=disable"),
		TripServiceAddr:     getEnv("TRIP_SERVICE_ADDR", "localhost:50051"),
		DriverServiceAddr:   getEnv("DRIVER_SERVICE_ADDR", "localhost:50052"),
		RiderServiceAddr:    getEnv("RIDER_SERVICE_ADDR", "localhost:50054"),
		MatchingServiceAddr: getEnv("MATCHING_SERVICE_ADDR", "localhost:50053"),
		Env:                 getEnv("ENV", "dev"),
	}
	log.Printf("config: env=%s port=%s", cfg.Env, cfg.Port)
	return cfg
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func mustDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		log.Fatalf("invalid duration for JWT_TTL: %v", err)
	}
	return d
}
