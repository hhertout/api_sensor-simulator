package core

import (
	"api_sensor/config"
	"api_sensor/core/entity"
	"api_sensor/core/tools"
	"api_sensor/schema"

	"github.com/mintance/go-uniqid"
)

func Main() {
	airSensorData := tools.GenerateAirSensorData()
	d := tools.GenerateFinalData(&airSensorData)

	err := saveInDatabase(&d)
	if err != nil {
		panic(err)
	}

}

func saveInDatabase(d *entity.AirData) error {
	id := uniqid.New(uniqid.Params{"", true})

	sensorData := d.GetData()
	result := config.DB.Create(&schema.AirData{
		ID:        id,
		Name:      d.GetName(),
		Location:  d.GetLocation(),
		TimeStamp: d.GetTimeStamp(),
		Data: schema.AirSensor{
			O2:        sensorData.GetO2(),
			Co2:       sensorData.GetCO2(),
			Co:        sensorData.GetCO(),
			So2:       sensorData.GetSO2(),
			Nox:       sensorData.GetNOX(),
			Pm10:      sensorData.GetPM10(),
			Pm25:      sensorData.GetPM25(),
			No2:       sensorData.GetNO2(),
			Ch4:       sensorData.GetCH4(),
			AirDataID: id,
		},
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
