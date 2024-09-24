package service

import (
	"fmt"
	"time"

	"github.com/herux/indegooweather/client"
	"github.com/herux/indegooweather/db"
	"github.com/herux/indegooweather/model"
)

func FetchWeather(apiKey string) error {
	var weatherResponse struct {
		Main struct {
			Temp     float64 `json:"temp"`
			Humidity int     `json:"humidity"`
		} `json:"main"`
		Wind struct {
			Speed float64 `json:"speed"`
		} `json:"wind"`
	}
	query := map[string]string{
		"q":     "Philadelphia",
		"appid": apiKey,
		"units": "metric",
	}

	client := client.New("https://api.openweathermap.org/data/2.5")
	_, err := client.GetJSON("/weather", &weatherResponse, query)
	if err != nil {
		fmt.Println("err: ", err)
		return err
	}

	weather := model.Weather{
		Temperature: weatherResponse.Main.Temp,
		Humidity:    weatherResponse.Main.Humidity,
		WindSpeed:   weatherResponse.Wind.Speed,
		Timestamp:   time.Now(),
	}

	db.DB.Create(&weather)
	return nil
}
