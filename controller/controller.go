package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/herux/indegooweather/service"
)

func HandleFetchIndego(c *fiber.Ctx) error {
	err := service.FetchAndStoreIndegoData()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to fetch and store data",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data fetched and stored succesfully",
	})
}
