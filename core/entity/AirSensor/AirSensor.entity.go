package AirSensor

import "time"

type AirSensor struct {
	name      string
	location  string
	timeStamp int64
	values    AirData
}

func New() AirSensor {
	values := AirData{
		o2:   210,
		co2:  300,
		co:   10,
		so2:  50,
		nox:  2,
		pm10: 30,
		pm25: 30,
		no2:  780,
		ch4:  1.72,
	}

	return AirSensor{
		name:      "Air Sensor",
		location:  "Jakarta",
		timeStamp: time.Now().Unix(),
		values:    values,
	}
}
func (d *AirSensor) GetName() string {
	return d.name
}
func (d *AirSensor) GetLocation() string {
	return d.location
}
func (d *AirSensor) GetTimeStamp() int64 {
	return d.timeStamp
}
func (d *AirSensor) GetData() AirData {
	return d.values
}
func (d *AirSensor) SetName(name string) {
	d.name = name
}
func (d *AirSensor) SetLocation(location string) {
	d.location = location
}
func (d *AirSensor) SetTimeStamp(timeStamp int64) {
	d.timeStamp = timeStamp
}
func (d *AirSensor) SetData(values AirData) {
	d.values = values
}
