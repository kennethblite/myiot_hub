package recommendation

import (
	"myiot_hub/sensor"
)

var rec = map[string]string{
	"OKAY":          "Good Dry, please leave in for a few more minutes",
	"PERFECT":       "Perfect Dry, enjoy your clothes",
	"BAD":           "The end humidity of the clothes is not low enough. Longer dry time",
	"CHECK ELEMENT": "Your heating element may be damaged, No High temps were recorded",
	"LINT FILTER":   "Your temperature is dangerously high, please clean lint filter",
	"OVER DRY":      "You have overdried your clothes, please consider doing a larger load",
}

func Recommend(data []sensor.SensorData, maxtemp float64) string {
	if maxtemp > 90 {
		return rec["LINT FILTER"]
	}
	if maxtemp < 50 {
		return rec["CHECK ELEMENT"]
	}

	midhumid := data[len(data)-30].Humidity()
	endhumid := data[len(data)-1].Humidity()
	if endhumid > 5 && endhumid < 8 {
		return rec["OKAY"]
	}

	if endhumid >= 8 {
		return rec["BAD"]
	}

	if midhumid < 5 {
		return rec["OVER DRY"]
	}
	return rec["PERFECT"]
}
