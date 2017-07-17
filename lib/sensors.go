package lib

import (
	"strings"

	"github.com/axnion/gosensors"
)

var gosensorsInit = gosensors.Init
var gosensorsCleanup = gosensors.Cleanup
var gosensorsGetDetectedChips = gosensors.GetDetectedChips

type Sensors struct {
	Temps []Sensor
	Fans  []Sensor
	Volt  []Sensor
}

type Sensor struct {
	Label string
	Value float64
}

// TODO: Fix spelling on Label
func GetSensors() Sensors {
	var sensors Sensors
	gosensorsInit()
	defer gosensorsCleanup()
	chips := gosensorsGetDetectedChips()

	for _, chip := range chips {
		features := chip.Features
		for _, feature := range features {
			if strings.Contains(feature.Name, "temp") {
				sensors.Temps = append(sensors.Temps, Sensor{feature.Lable, feature.Value})
			}

			if strings.Contains(feature.Name, "fan") {
				sensors.Fans = append(sensors.Fans, Sensor{feature.Lable, feature.Value})
			}

			if strings.Contains(feature.Name, "in") {
				sensors.Volt = append(sensors.Volt, Sensor{feature.Lable, feature.Value})
			}
		}
	}

	return sensors
}
