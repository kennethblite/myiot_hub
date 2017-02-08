package sensor

import(
	"testing"
	"bufio"
	"os"
)


func TestSensorParsing(t *testing.T){
	file, err := os.Open("test_data2.txt")
	if (err != nil){
		t.Fatal(err)
	}
	defer file.Close()
	Reader := bufio.NewReader(file)
	for {
	data, err := Reader.ReadString(10)
	if err != nil {
		t.Fatal(err)
	}
	sensordata, new := ParseAdditionalData(data)
	if new {
		t.Log(sensordata)
	}
	}
}
	
