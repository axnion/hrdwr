package hrdwr

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/axnion/gosensors"
)

/*
 * - gosensors mocking -------------------------------------------------------------------------------------------------
 */

var gosensorsInitCallCounter = 0
var gosensorsCleanupCallCounter = 0
var gosensorsGetDetectedChipsCallCounter = 0

func gosensorsInitMock() {
	gosensorsInitCallCounter++
}

func gosensorsCleanupMock() {
	gosensorsCleanupCallCounter++
}

func gosensorsGetDetectedChipsMock() []gosensors.Chip {
	gosensorsGetDetectedChipsCallCounter++
	return createChipMocks()
}

func prepareGosensorMocks() {
	gosensorsInit = gosensorsInitMock
	gosensorsCleanup = gosensorsCleanupMock
	gosensorsGetDetectedChips = gosensorsGetDetectedChipsMock
}

func createChipMocks() []gosensors.Chip {
	var chips []gosensors.Chip
	acpi := gosensors.Chip{
		Prefix: 	"atk0110",
		Path:		"/sys/class/hwmon/hwmon0",
		Addr: 		0,
		AdapterName:	"ACPI interface",
		Features: createACPIFeatures(),
	}
	chips = append(chips, acpi)

	isa := gosensors.Chip {
		Prefix: 		"coretemp",
		Path: 			"/sys/class/hwmon/hwmon1",
		Addr:			0,
		AdapterName:	"ISA adapter",
		Features: 		createISAFeatures(),
	}
	chips = append(chips, isa)

	return chips
}

func createACPIFeatures() []gosensors.Feature {
	var features []gosensors.Feature

	in0 := gosensors.Feature {
		Name: "in0",
		Lable: "VCore Voltage",
		Type: 0,
		Number: 0,
		Value: 0.944,
	}
	features = append(features, in0)

	in1 := gosensors.Feature {
		Name: "in1",
		Lable: " +3.3 Voltage",
		Type: 0,
		Number: 1,
		Value: 3.28,
	}
	features = append(features, in1)

	fan1 := gosensors.Feature {
		Name: "fan1",
		Lable: "CPU FAN Speed",
		Type: 1,
		Number: 2,
		Value: 2481,
	}
	features = append(features, fan1)

	fan2 := gosensors.Feature {
		Name: "fan2",
		Lable: "CHASSI FAN Speed",
		Type: 1,
		Number: 3,
		Value: 0,
	}
	features = append(features, fan2)

	temp1 := gosensors.Feature {
		Name: "temp1",
		Lable: "CPU Temperature",
		Type: 2,
		Number: 4,
		Value: 41,
	}
	features = append(features, temp1)

	temp2 := gosensors.Feature {
		Name: "temp2",
		Lable: "MB Temperature",
		Type: 2,
		Number: 5,
		Value: 37,
	}
	features = append(features, temp2)

	return features
}

func createISAFeatures() []gosensors.Feature {
	var features []gosensors.Feature

	temp2 := gosensors.Feature {
		Name: "temp2",
		Lable: "Core 0",
		Type: 2,
		Number: 0,
		Value: 46,
	}
	features = append(features, temp2)

	temp3 := gosensors.Feature {
		Name: "temp3",
		Lable: "Core 1",
		Type: 2,
		Number: 1,
		Value: 44,
	}
	features = append(features, temp3)

	return features
}

/*
 * - Test Suite ------------------------------------------------------------------------------------
 */

func TestGetSensors(t *testing.T) {
	prepareGosensorMocks()

	results := GetSensors()

	assert.Equal(t, 1, gosensorsInitCallCounter)
	assert.Equal(t, 1, gosensorsCleanupCallCounter)
	assert.Equal(t, 1, gosensorsGetDetectedChipsCallCounter)

	assert.Equal(t, 4, len(results.Temps))
	assert.Equal(t, 2, len(results.Fans))
	assert.Equal(t, 2, len(results.Volt))

	// Checking Temps
	assert.Equal(t, "CPU Temperature", results.Temps[0].Label)
	assert.Equal(t, float64(41), results.Temps[0].Value)
	assert.Equal(t, "MB Temperature", results.Temps[1].Label)
	assert.Equal(t, float64(37), results.Temps[1].Value)
	assert.Equal(t, "Core 0", results.Temps[2].Label)
	assert.Equal(t, float64(46), results.Temps[2].Value)
	assert.Equal(t, "Core 1", results.Temps[3].Label)
	assert.Equal(t, float64(44), results.Temps[3].Value)

	// Checking Fans
	assert.Equal(t, "CPU FAN Speed", results.Fans[0].Label)
	assert.Equal(t, float64(2481), results.Fans[0].Value)
	assert.Equal(t, "CHASSI FAN Speed", results.Fans[1].Label)
	assert.Equal(t, float64(0), results.Fans[1].Value)

	// Checking Voltages
	assert.Equal(t, "VCore Voltage", results.Volt[0].Label)
	assert.Equal(t, float64(0.944), results.Volt[0].Value)
	assert.Equal(t, " +3.3 Voltage", results.Volt[1].Label)
	assert.Equal(t, float64(3.28), results.Volt[1].Value)
}
