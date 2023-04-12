package schema

import "gorm.io/gorm"

type AirData struct {
	gorm.Model
	ID        string    `gorm:"primaryKey, unique`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	TimeStamp int64     `json:"timestamp"`
	Data      AirSensor `gorm:"foreignKey:AirDataID"`
}
