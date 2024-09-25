package route

import (
	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
	"github.com/herux/indegooweather/controller"
	_ "github.com/herux/indegooweather/swagger/docs"
)

func RegisterAPI(route fiber.Router) {
	route.Get("/swagger/*", swagger.HandlerDefault)

	apiv1 := route.Group("/api/v1")
	apiv1.Post("/indego-data-fetch-and-store-it-db", controller.HandleFetchIndego)
	apiv1.Get("/stations", controller.HandleStationSnapshot)
	apiv1.Get("/stations/:kioskId", controller.HandleKiosk)
}
