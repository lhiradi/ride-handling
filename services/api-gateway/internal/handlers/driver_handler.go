package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lhiradi/ride-handling/services/api-gateway/internal/services"
)

func SetStatusHandler(svc *services.DriverService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		driverID := c.Params("id")
		var req struct {
			Status string `json:"status"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
		}
		if driverID == "" || req.Status == "" {
			return c.Status(400).JSON(fiber.Map{"error": "driver_id and status required"})
		}
		if err := svc.SetStatus(c.Context(), driverID, req.Status); err != nil {
			return c.Status(502).JSON(fiber.Map{"error": err.Error()})
		}
		return c.SendStatus(204)
	}
}

func HeartbeatHandler(svc *services.DriverService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		driverID := c.Params("id")
		var req struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
		}
		if driverID == "" {
			return c.Status(400).JSON(fiber.Map{"error": "driver_id required"})
		}
		ts := time.Now().Unix()
		if err := svc.Heartbeat(c.Context(), driverID, req.Lat, req.Lon, ts); err != nil {
			return c.Status(502).JSON(fiber.Map{"error": err.Error()})
		}
		return c.SendStatus(204)
	}
}
