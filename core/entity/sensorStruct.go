package entity

type AirSensor struct {
	o2   float32
	co2  float32
	co   float32
	so2  float32
	nox  float32
	pm10 float32
	pm25 float32
	no2  float32
	ch4  float32
}

func (d *AirSensor) GetO2() float32 {
	return d.o2
}
func (d *AirSensor) GetCO2() float32 {
	return d.co2
}
func (d *AirSensor) GetCO() float32 {
	return d.co
}
func (d *AirSensor) GetSO2() float32 {
	return d.so2
}
func (d *AirSensor) GetNOX() float32 {
	return d.nox
}
func (d *AirSensor) GetPM10() float32 {
	return d.pm10
}
func (d *AirSensor) GetPM25() float32 {
	return d.pm25
}
func (d *AirSensor) GetNO2() float32 {
	return d.no2
}
func (d *AirSensor) GetCH4() float32 {
	return d.ch4
}
func (d *AirSensor) SetO2(o2 float32) {
	d.o2 = o2
}
func (d *AirSensor) SetCO2(co2 float32) {
	d.co2 = co2
}
func (d *AirSensor) SetCO(co float32) {
	d.co = co
}
func (d *AirSensor) SetSO2(so2 float32) {
	d.so2 = so2
}
func (d *AirSensor) SetNOX(nox float32) {
	d.nox = nox
}
func (d *AirSensor) SetPM10(pm10 float32) {
	d.pm10 = pm10
}
func (d *AirSensor) SetPM25(pm25 float32) {
	d.pm25 = pm25
}
func (d *AirSensor) SetNO2(no2 float32) {
	d.no2 = no2
}
func (d *AirSensor) SetCH4(ch4 float32) {
	d.ch4 = ch4
}
