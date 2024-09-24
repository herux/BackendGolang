package model

import (
	"time"

	"gorm.io/gorm"
)

type BikeStation struct {
	gorm.Model
	StationID      int
	AvailableBikes int
	AvailableDocks int
	TotalDocks     int
	Timestamp      time.Time
}
