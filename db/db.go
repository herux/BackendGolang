package db

import (
	"log"
	"strconv"

	"github.com/herux/indegooweather/config"
	"github.com/herux/indegooweather/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(isDryRun bool) {
	dsn := "host=" + config.DatabaseConfig().Host + " user=" + config.DatabaseConfig().User +
		" password=" + config.DatabaseConfig().Password + " dbname=" + config.DatabaseConfig().Dbname +
		" port=" + strconv.FormatUint(uint64(config.DatabaseConfig().Port), 10) +
		" sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DryRun: isDryRun,
	})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	DB.AutoMigrate(&model.BikeStation{}, &model.Weather{})
}
