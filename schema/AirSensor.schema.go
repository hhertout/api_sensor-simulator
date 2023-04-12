package schema

type AirSensor struct {
	ID        uint    `gorm:"primaryKey"`
	AirDataID string  `json:"air_data_id"`
	O2        float32 `json:"o2"`
	Co2       float32 `json:"co2"`
	Co        float32 `json:"co"`
	So2       float32 `json:"so2"`
	Nox       float32 `json:"nox"`
	Pm10      float32 `json:"pm10"`
	Pm25      float32 `json:"pm25"`
	No2       float32 `json:"no2"`
	Ch4       float32 `json:"ch4"`
}
