package entity

type AirData struct {
	name      string
	location  string
	timeStamp int64
	data      AirSensor
}

func (d *AirData) GetName() string {
	return d.name
}
func (d *AirData) GetLocation() string {
	return d.location
}
func (d *AirData) GetTimeStamp() int64 {
	return d.timeStamp
}
func (d *AirData) GetData() AirSensor {
	return d.data
}
func (d *AirData) SetName(name string) {
	d.name = name
}
func (d *AirData) SetLocation(location string) {
	d.location = location
}
func (d *AirData) SetTimeStamp(timeStamp int64) {
	d.timeStamp = timeStamp
}
func (d *AirData) SetData(data AirSensor) {
	d.data = data
}
