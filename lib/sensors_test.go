package lib

import (
	"testing"
	"github.com/md14454/gosensors"
	"github.com/stretchr/testify/assert"
)

/*
 * - GoSensors Mocking -------------------------------------------------------------------------------------------------
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
	var chips []gosensors.Chip

	gosensorsGetDetectedChipsCallCounter++
	return chips
}

func prepareGosensorMocks() {
	gosensorsInit = gosensorsInitMock
	gosensorsCleanup = gosensorsCleanupMock
	gosensorsGetDetectedChips = gosensorsGetDetectedChipsMock
}

/*
 * - Test Suite ------------------------------------------------------------------------------------
 */

func TestGetSensors(t *testing.T) {
	prepareGosensorMocks()

	GetSensors()

	assert.Equal(t, 1, gosensorsInitCallCounter)
	assert.Equal(t, 1, gosensorsCleanupCallCounter)
	assert.Equal(t, 1, gosensorsGetDetectedChipsCallCounter)
}
