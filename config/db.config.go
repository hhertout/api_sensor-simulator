package config

import (
	"api_sensor/schema"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconnect() {
	var err error
	dsn := os.Getenv("DATABASE")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func Migrate() {
	DB.AutoMigrate(&schema.Auth{})
	DB.AutoMigrate(&schema.AirSensor{})
	DB.AutoMigrate(&schema.AirData{})
}
