package AirSensor

type AirData struct {
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

func (d *AirData) GetO2() float32 {
	return d.o2
}
func (d *AirData) GetCO2() float32 {
	return d.co2
}
func (d *AirData) GetCO() float32 {
	return d.co
}
func (d *AirData) GetSO2() float32 {
	return d.so2
}
func (d *AirData) GetNOX() float32 {
	return d.nox
}
func (d *AirData) GetPM10() float32 {
	return d.pm10
}
func (d *AirData) GetPM25() float32 {
	return d.pm25
}
func (d *AirData) GetNO2() float32 {
	return d.no2
}
func (d *AirData) GetCH4() float32 {
	return d.ch4
}
func (d *AirData) SetO2(o2 float32) {
	d.o2 = o2
}
func (d *AirData) SetCO2(co2 float32) {
	d.co2 = co2
}
func (d *AirData) SetCO(co float32) {
	d.co = co
}
func (d *AirData) SetSO2(so2 float32) {
	d.so2 = so2
}
func (d *AirData) SetNOX(nox float32) {
	d.nox = nox
}
func (d *AirData) SetPM10(pm10 float32) {
	d.pm10 = pm10
}
func (d *AirData) SetPM25(pm25 float32) {
	d.pm25 = pm25
}
func (d *AirData) SetNO2(no2 float32) {
	d.no2 = no2
}
func (d *AirData) SetCH4(ch4 float32) {
	d.ch4 = ch4
}
