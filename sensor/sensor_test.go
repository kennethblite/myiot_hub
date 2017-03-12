package sensor

import (
	"bufio"
	"os"
	"testing"
)

func TestSensorParsing(t *testing.T) {
	file, err := os.Open("large.txt")
	Buffer := NewRotateBuffer(10)
	if err != nil {
		t.Fatal(err)
	}
	var steadystate bool
	defer file.Close()
	Reader := bufio.NewReader(file)
	for {
		data, err := Reader.ReadString(10)
		if err != nil {
			t.Fatal(err)
		}
		sensordata, new := ParseAdditionalData(data)
		if new {
			Buffer.Add(sensordata.temp)
			if Buffer.SlopeChange() {
				if Buffer.Average() > 80 || steadystate {
					steadystate = true
					t.Log(sensordata.timing)
					t.Log(Buffer.Average())
					t.Log(Buffer.SampleNumber())
				}
			}
		}
	}
}
