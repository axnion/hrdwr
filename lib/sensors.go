package lib

import (
	"github.com/md14454/gosensors"
	"strings"
)

var gosensorsInit = gosensors.Init
var gosensorsCleanup = gosensors.Cleanup
var gosensorsGetDetectedChips = gosensors.GetDetectedChips

type SensorMon struct {
}

type SensorLib interface {
	Init()
	Cleanup()
	GetDetectedChips()
}

type Sensors struct {
	Temps []Sensor
	Fans []Sensor
	Volt []Sensor
}

type Sensor struct {
	Label string
	Value float64
}

func NewSensorMon() SensorMon {
	return SensorMon {}
}

func (mon SensorMon) GetSensors() Sensors {
	return getSensors()
}

func getSensors() Sensors {
	var sensors Sensors
	gosensorsInit()
	defer gosensorsCleanup()
	chips := gosensorsGetDetectedChips()

	for _, chip := range chips {
		features := chip.GetFeatures()
		for _, feature := range features {
			if strings.Contains(feature.Name, "temp") {
				sensors.Temps = append(sensors.Temps, Sensor{feature.GetLabel(), feature.GetValue()})
			}

			if strings.Contains(feature.Name, "fan") {
				sensors.Fans = append(sensors.Fans, Sensor{feature.GetLabel(), feature.GetValue()})
			}

			if strings.Contains(feature.Name, "in") {
				sensors.Volt = append(sensors.Volt, Sensor{feature.GetLabel(), feature.GetValue()})
			}
		}
	}

	return sensors
}
