package tools

import (
	"api_sensor/core/entity"
	"math/rand"
)

func GenerateAirSensorData() entity.AirSensor {
	sensor := createDataSet()

	sensorData := genRandomDataFrom(sensor)

	return sensorData
}

func GenerateFinalData(s *entity.AirSensor) entity.AirData {
	var data entity.AirData
	ts := GetTimestamp()

	data.SetName("Air Sensor")
	data.SetLocation("Jakarta")
	data.SetTimeStamp(ts)
	data.SetData(*s)

	return data
}

func createDataSet() entity.AirSensor {
	var defaultValues entity.AirSensor

	defaultValues.SetO2(210)
	defaultValues.SetCO2(300)
	defaultValues.SetCO(10)
	defaultValues.SetSO2(50)
	defaultValues.SetNOX(2)
	defaultValues.SetPM10(30)
	defaultValues.SetPM25(30)
	defaultValues.SetNO2(780)
	defaultValues.SetCH4(1.72)

	return defaultValues
}

func genRandomDataFrom(d entity.AirSensor) entity.AirSensor {
	d.SetO2(randomiser(d.GetO2()))
	d.SetCO2(randomiser(d.GetCO2()))
	d.SetCO(randomiser(d.GetCO()))
	d.SetSO2(randomiser(d.GetSO2()))
	d.SetNOX(randomiser(d.GetNOX()))
	d.SetPM10(randomiser(d.GetPM10()))
	d.SetPM25(randomiser(d.GetPM25()))
	d.SetNO2(randomiser(d.GetNO2()))
	d.SetCH4(randomiser(d.GetCH4()))

	return d
}

func randomiser(data float32) float32 {
	return data + (data * (rand.Float32() - 0.5))
}
