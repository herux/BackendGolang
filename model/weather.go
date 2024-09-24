package model

import (
	"time"

	"gorm.io/gorm"
)

type Weather struct {
	gorm.Model
	Temperature float64
	Humidity    int
	WindSpeed   float64
	Timestamp   time.Time
}
