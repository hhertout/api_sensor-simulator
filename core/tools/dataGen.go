package tools

import (
	"api_sensor/core/entity/AirSensor"
	"math/rand"
)

func GenerateAirSensorData() AirSensor.AirData {
	sensor := AirSensor.New()

	sensorData := genRandomDataFrom(sensor.GetData())

	return sensorData
}

func GenerateFinalData(s AirSensor.AirData) AirSensor.AirSensor {
	var data AirSensor.AirSensor
	ts := GetTimestamp()

	data.SetName("Air Sensor")
	data.SetLocation("Jakarta")
	data.SetTimeStamp(ts)
	data.SetData(s)

	return data
}

func genRandomDataFrom(d AirSensor.AirData) AirSensor.AirData {
	d.SetO2(Randomizer(d.GetO2()))
	d.SetCO2(Randomizer(d.GetCO2()))
	d.SetCO(Randomizer(d.GetCO()))
	d.SetSO2(Randomizer(d.GetSO2()))
	d.SetNOX(Randomizer(d.GetNOX()))
	d.SetPM10(Randomizer(d.GetPM10()))
	d.SetPM25(Randomizer(d.GetPM25()))
	d.SetNO2(Randomizer(d.GetNO2()))
	d.SetCH4(Randomizer(d.GetCH4()))

	return d
}

func Randomizer(data float32) float32 {
	return data + (data * (rand.Float32() - 0.5))
}
