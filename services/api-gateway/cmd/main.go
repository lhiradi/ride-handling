package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	commonv1 "github.com/lhiradi/ride-handling/proto/common/v1"
	riderv1 "github.com/lhiradi/ride-handling/proto/rider/v1"
	tripv1 "github.com/lhiradi/ride-handling/proto/trip/v1"
	"github.com/lhiradi/ride-handling/services/api-gateway/internal/clients"
	"google.golang.org/grpc"
)

func main() {
	app := fiber.New()

	// gRPC connections
	riderConn, _ := grpc.Dial("localhost:50054", grpc.WithInsecure())
	tripConn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())

	riderClient := clients.NewRiderClient(riderConn)
	tripClient := tripv1.NewTripServiceClient(tripConn)

	// Create Rider
	app.Post("/riders", func(c *fiber.Ctx) error {
		var req struct {
			Name     string `json:"name"`
			Phone    string `json:"phone"`
			Email    string `json:"email"`
			Language string `json:"language"`
		}
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}
		resp, err := riderClient.CreateRider(c.Context(), &riderv1.CreateRiderRequest{
			Name: req.Name, Phone: req.Phone, Email: req.Email, Language: req.Language,
		})
		if err != nil {
			return fiber.NewError(502, err.Error())
		}
		return c.JSON(fiber.Map{"rider_id": resp.RiderId})
	})

	// Create Trip
	app.Post("/trips", func(c *fiber.Ctx) error {
		var req struct {
			RiderID string            `json:"rider_id"`
			Pickup  commonv1.GeoPoint `json:"pickup"`
			Dropoff commonv1.GeoPoint `json:"dropoff"`
		}
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}
		resp, err := tripClient.CreateTrip(c.Context(), &tripv1.CreateTripRequest{
			RiderId: req.RiderID,
			Pickup:  &req.Pickup,
			Dropoff: &req.Dropoff,
		})
		if err != nil {
			return fiber.NewError(502, err.Error())
		}
		return c.JSON(resp.Trip)
	})

	log.Println("API Gateway listening on :8080")
	log.Fatal(app.Listen(":8080"))
}
