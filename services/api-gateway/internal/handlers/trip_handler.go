package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lhiradi/ride-handling/services/api-gateway/internal/services"
)

func CreateTripHandler(svc *services.TripService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			RiderID   string  `json:"rider_id"`
			PickupLat float64 `json:"pickup_lat"`
			PickupLon float64 `json:"pickup_lon"`
			DropLat   float64 `json:"drop_lat"`
			DropLon   float64 `json:"drop_lon"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
		}

		resp, err := svc.CreateTrip(c.Context(), req.RiderID, req.PickupLat, req.PickupLon, req.DropLat, req.DropLon)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(resp)
	}
}
