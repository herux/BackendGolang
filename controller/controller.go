package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/herux/indegooweather/db"
	"github.com/herux/indegooweather/model"
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

func HandleStationSnapshot(c *fiber.Ctx) error {
	at := c.Query("at")
	t, err := time.Parse(time.RFC3339, at)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid timestamp format",
		})
	}

	var stations []model.BikeStation
	db.DB.Where("timestamp >= ?", t).Order("timestamp asc").Find(&stations)
	if len(stations) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No data available at the specified time",
		})
	}
	var weather model.Weather
	db.DB.Where("timestamp >= ?", t).Order("timestamp asc").First(&weather)

	return c.JSON(fiber.Map{
		"at":       stations[0].Timestamp,
		"stations": stations,
		"weather":  weather,
	})
}

func HandleKiosk(c *fiber.Ctx) error {
	kioskId := c.Params("kioskId")
	at := c.Query("at")
	t, err := time.Parse(time.RFC3339, at)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid timestamp format",
		})
	}

	var station model.BikeStation
	db.DB.Where("station_id = ? AND timestamp >= ?", kioskId, t).Order("timestamp asc").First(&station)
	if station.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No data available at the specified time for this station",
		})
	}

	var weather model.Weather
	db.DB.Where("timestamp >= ?", at).Order("timestamp asc").First(&weather)

	return c.JSON(fiber.Map{
		"at":      station.Timestamp,
		"station": station,
		"weather": weather,
	})
}
