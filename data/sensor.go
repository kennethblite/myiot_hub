package sensor


import(
	"strings"
	"strconv"
	"fmt"
)


var CurrentData dataStack

type dataStack struct{
	data string 
}

type SensorData struct{
 accelx float64 
 accely float64
 accelz float64
 temp float64
 humidity float64
}


func ParseAdditionalData(s string) (SensorData, bool){
	CurrentData.data += s
	if !strings.Contains(CurrentData.data, "}"){
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
	s.accelx,_ = strconv.ParseFloat(splitdata[0],64)	
	s.accely,_ = strconv.ParseFloat(splitdata[1],64)	
	s.accelz,_ = strconv.ParseFloat(splitdata[2],64)
	twopoints := strings.Split(splitdata[3], "C")
	fmt.Println(twopoints)
	s.temp, _ = strconv.ParseFloat(twopoints[0],64)
	s.humidity, _ = strconv.ParseFloat(twopoints[1],64)
	return s	
}




