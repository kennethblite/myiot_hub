package web

import (
	"myiot_hub/sensor"
	"testing"
)

func TestPostPoint(t *testing.T) {
	t.Fatal(PostPoint(sensor.SensorData{}, "1340"))
}

func TestCreateCycle(t *testing.T) {
	t.Fatal(CreateDryingCycle("1340"))
}

func TestFinishDryingCycle(t *testing.T) {
	t.Fatal(FinishDryingCycle("1340", "Great job"))
}
