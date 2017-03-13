package sensor

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"encoding/hex"
)

var CurrentData dataStack

type dataStack struct {
	data string
}

type DryerCycle struct {
	Data []SensorData
}

func NewDryerCycle()DryerCycle{
	data := make([]SensorData, 1, 1)
	return DryerCycle{data}
}

type SensorData struct {
	accelx   float64
	accely   float64
	accelz   float64
	temp     float64
	humidity float64
	timing   time.Time
}

func (s SensorData) Humidity() float64 {
	return s.humidity
}
func (s SensorData) Temp() float64 {
	return s.temp
}

func ParseAdditionalData(s string) (SensorData, bool) {
	CurrentData.data += s
	if !strings.Contains(CurrentData.data, "}") {
		return SensorData{}, false
	}
	NewSensorData := strings.Split(CurrentData.data, "}")
	CurrentData.data = NewSensorData[1]
	return ParseSensor(NewSensorData[0]), true
}

func ParseSensor(sensordata string) SensorData {
	s := SensorData{}
	sensordata = strings.Replace(sensordata, " ", "", -1)
	sensordata = strings.Replace(sensordata, "%", "", -1)
	sensordata = strings.Replace(sensordata, "{", "", -1)
	sensordata = strings.Replace(sensordata, "\n", "", -1)
	sensordata = strings.Replace(sensordata, ";", "", -1)
	splitdata := strings.Split(sensordata, ",")
	fmt.Println(sensordata)
	//s.accelx, _ = strconv.ParseFloat(splitdata[0], 64)
	//s.accely, _ = strconv.ParseFloat(splitdata[1], 64)
	//s.accelz, _ = strconv.ParseFloat(splitdata[2], 64)
	twopoints := strings.Split(splitdata[0], "C")
	fmt.Println(twopoints)
	s.temp, _ = strconv.ParseFloat(twopoints[0], 64)
	s.humidity, _ = strconv.ParseFloat(twopoints[1], 64)
	s.timing = time.Now()
	return s
}

func TrimNonsense(s string)string{
	s = strings.Replace(s, "Notification handle = 0x0012 value: ","",-1)
	s = strings.Replace(s, " ", "",-1)
	s = strings.Replace(s, "\n", "",-1)
	fmt.Println(s)
	newstring,err := hex.DecodeString(s)
	if err != nil{
		fmt.Println(err)
	}	
	return string(newstring)
}
