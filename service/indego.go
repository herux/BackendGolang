package service

import (
	"log"
	"time"

	"github.com/herux/indegooweather/client"
	"github.com/herux/indegooweather/db"
	"github.com/herux/indegooweather/model"
)

func FetchAndStoreIndegoData() error {

	var stationResponse struct {
		Features []struct {
			Properties struct {
				ID             int `json:"kioskId"`
				BikesAvailable int `json:"bikesAvailable"`
				DocksAvailable int `json:"docksAvailable"`
				TotalDocks     int `json:"totalDocks"`
			} `json:"properties"`
		} `json:"features"`
	}

	client := client.New("https://www.rideindego.com")
	_, err := client.GetJSON("/stations/json/", &stationResponse, nil)
	if err != nil {
		log.Println("Error fetching bike station data:", err)
		return err
	}

	for _, feature := range stationResponse.Features {
		station := model.BikeStation{
			StationID:      feature.Properties.ID,
			AvailableBikes: feature.Properties.BikesAvailable,
			AvailableDocks: feature.Properties.DocksAvailable,
			TotalDocks:     feature.Properties.TotalDocks,
			Timestamp:      time.Now(),
		}
		db.DB.Create(&station)
	}
	return nil
}
