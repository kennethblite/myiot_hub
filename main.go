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
)

var userstream DataStream
var Buffer *sensor.RotateBuffer
var steadystate bool
var cycle DryerCycle
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
	setup()
	cmd := exec.Command("gatttool", "-b", "20:91:48:4D:94:BC", "--char-write-req", "-a", "0x0013", "-n", "0100", "--listen")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}
	Reader := bufio.NewReader(stdout)
	Buffer = sensor.NewRotateBuffer(10)
	for {
		Message, err := Reader.ReadString(10)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("device Or Resource Busy")
				os.Exit(1)
			}
			fmt.Println(err)
		}
		if !Drying {
			Drying = true
			web.CreateDryingCycle(fmt.Sprint(tag))
			cycle = sensor.NewDryerCycle()
		}
		go ProcessReadings(Message)
	}
}

func setup() {
	tag = 40
	Max = 0
}

func ProcessReadings(data string) {
	for {
		sensordata, new := sensor.ParseAdditionalData(data)
		if new {
			web.PostPoint(sensordata)
			cycle.Data = append(cycle.Data, sensordata)
			Buffer.Add(sensordata.Temp())
			if (Max < sensordata.Temp()) {
				Max = sensordata.Temp()
			}
			if Buffer.SlopeChange() {
				if Buffer.Average() > 80 || steadystate {
					steadystate = true
					fmt.Println(Buffer.Average())
					fmt.Println(Buffer.SampleNumber())
				}
				if steadystate && Buffer.Average() < 55 {
					drying = false
					message := recommendation.Recommend(cycle.Data)
					web.FinishDryingCycle(fmt.Sprint(tag), message, Max)
					os.Exit(1)
				}
			}
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
}
