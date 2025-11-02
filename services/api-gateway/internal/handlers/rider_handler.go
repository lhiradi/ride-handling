package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lhiradi/ride-handling/services/api-gateway/internal/services"
)

func CreateRiderHandler(svc *services.RiderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			Name     string `json:"name"`
			Phone    string `json:"phone"`
			Email    string `json:"email"`
			Language string `json:"language"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
		}
		if req.Name == "" || req.Phone == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "name and phone are required"})
		}

		resp, err := svc.CreateRider(c.Context(), req.Name, req.Phone, req.Email, req.Language)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"rider_id": resp.RiderId})
	}
}

func GetRiderHandler(svc *services.RiderService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		riderID := c.Params("id")
		if riderID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "rider id is required"})
		}

		resp, err := svc.GetRider(c.Context(), riderID)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"name":     resp.Name,
			"phone":    resp.Phone,
			"email":    resp.Email,
			"language": resp.Language,
		})
	}
}
