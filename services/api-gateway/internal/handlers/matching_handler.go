package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lhiradi/ride-handling/services/api-gateway/internal/services"
)

func MatchHandler(svc *services.MatchingService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			TripID string  `json:"trip_id"`
			Lat    float64 `json:"lat"`
			Lon    float64 `json:"lon"`
			Radius float64 `json:"radius_km"`
			Limit  int32   `json:"limit"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
		}

		resp, err := svc.MatchTrip(c.Context(), req.TripID, req.Lat, req.Lon, req.Radius, req.Limit)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(resp)
	}
}
func AcceptInvitationHandler(svc *services.MatchingService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			InvitationID string `json:"invitation_id"`
			TripID       string `json:"trip_id"`
			DriverID     string `json:"driver_id"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
		}
		if req.InvitationID == "" || req.TripID == "" || req.DriverID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invitation_id, trip_id, and driver_id are required"})
		}

		resp, err := svc.AcceptInvitation(c.Context(), req.InvitationID, req.TripID, req.DriverID)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"accepted": resp.Accepted})
	}
}
