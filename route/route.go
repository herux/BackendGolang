package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/herux/indegooweather/controller"
)

func RegisterAPI(route fiber.Router) {
	apiv1 := route.Group("/api/v1")
	apiv1.Post("/indego-data-fetch-and-store-it-db", controller.HandleFetchIndego)
}
