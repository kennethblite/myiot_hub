package main

import (
	"fmt"
	"os"
	"os/exec"
	//"strings"
	"bufio"
	//"encoding/hex"
	"myiot_hub/sensor"
	"myiot_hub/sensor/recommendation"
	"myiot_hub/web"
	"time"
	"strings"
)

var userstream DataStream
var Buffer *sensor.RotateBuffer
var steadystate bool
var cycle sensor.DryerCycle
var tag int
var Drying bool
var Max float64

type DataStream struct {
	//s []Stats
	username string
	dryer    string
}

//gatttool -b 20:91:48:4D:94:BC --char-write-req -a 0x0013 -n 0100 --listen

func main() {
	for {
		setup()
	}
}

func setup() {
	tag = 45
	Max = 0
	cmd := exec.Command("gatttool", "-b", "20:91:48:4D:94:BC", "--char-write-req", "-a", "0x0013", "-n", "0100", "--listen")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("beginning setup")
	/*file, _ := os.Open("sensor/large.txt")
	defer file.Close()
	*/
	Reader := bufio.NewReader(stdout)
	Buffer = sensor.NewRotateBuffer(10)
	for {
		s := make(chan bool, 1)
		go func() {
			Message, err := Reader.ReadString(10)
			if err != nil {
				if err.Error() == "EOF" {
					fmt.Println("Please Check Sensor Connections")
					web.FinishDryingCycle(fmt.Sprint(tag), "Please Check Dryer Connections")
					os.Exit(1)
				}
				fmt.Println(err)
			}
			if !Drying {
				Drying = true
				web.CreateDryingCycle(fmt.Sprint(tag))
				cycle = sensor.NewDryerCycle()
			}
			if ProcessReadings(Message) {
				fmt.Println("Drying finished")
				os.Exit(1)
			}
			s <- false
		}()
		go func() {
			time.Sleep(20 * time.Second)
			s <- true
		}()
		if <-s {
			fmt.Println("Error communicating, retrying")
			break
		}
	}
}

func ProcessReadings(data string) bool {
		if strings.Contains(data, "characteristic"){
			return false
		}
		data = sensor.TrimNonsense(data)
		fmt.Println(data)
		sensordata, new := sensor.ParseAdditionalData(data)
		if new {
			if sensordata.Temp() > 200 {
				return false 
			}
			if sensordata.Humidity() < 1 {
				return false
			}
			fmt.Println(sensordata)
			web.PostPoint(sensordata, fmt.Sprint(tag))
			cycle.Data = append(cycle.Data, sensordata)
			Buffer.Add(sensordata.Temp())
			if Max < sensordata.Temp() {
				Max = sensordata.Temp()
			}
			if Buffer.SlopeChange() {
				if Buffer.Average() > 80 || steadystate {
					steadystate = true
					fmt.Println(Buffer.Average())
					fmt.Println(Buffer.SampleNumber())
				}
			}
			
				if steadystate && sensordata.Temp() < 55 {
					Drying = false
					message := recommendation.Recommend(cycle.Data, Max)
					web.FinishDryingCycle(fmt.Sprint(tag), message)
					time.Sleep(10*time.Second)
					return true
				}
		}
	/*
		final_string := ""
		s = strings.Replace(s, "\n","",-1)
		s = strings.Replace(s, "Notification handle = 0x0012 value: ", "",-1)
		s = strings.Replace(s, " " , "",-1)
		s = strings.Replace(s, ";" , "",-1)
		final_string = s

		//fmt.Println(final_string)
		decoded, _  := hex.DecodeString(final_string)
		fmt.Println(string(decoded))*/
		return false
}
