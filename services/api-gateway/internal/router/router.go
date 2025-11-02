package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lhiradi/ride-handling/services/api-gateway/internal/handlers"
	"github.com/lhiradi/ride-handling/services/api-gateway/internal/services"
)

func Setup(
	app *fiber.App,
	driverSvc *services.DriverService,
	riderSvc *services.RiderService,
	tripSvc *services.TripService,
	matchingSvc *services.MatchingService,
) {
	// Driver routes
	app.Post("/drivers/:id/status", handlers.SetStatusHandler(driverSvc))
	app.Post("/drivers/:id/heartbeat", handlers.HeartbeatHandler(driverSvc))

	// Rider routes
	app.Post("/riders", handlers.CreateRiderHandler(riderSvc))
	app.Get("/riders/:id", handlers.GetRiderHandler(riderSvc))

	// Trip routes
	app.Post("/trips", handlers.CreateTripHandler(tripSvc))

	// Matching routes
	app.Post("/match", handlers.MatchHandler(matchingSvc))
	app.Post("/accept", handlers.AcceptInvitationHandler(matchingSvc))
}
