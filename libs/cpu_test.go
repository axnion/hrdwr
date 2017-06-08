package libs

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strconv"
	"github.com/axnion/hrdwr/libs/parse"
)

/*
 * - Test Runner -----------------------------------------------------------------------------------
 */

type TestRunner struct {
	executions int
}

func (runner TestRunner) Run(cmd string, args ...string) ([]byte, error) {
	runner.executions++
	return []byte("/proc/stat content"), nil
}

/*
 * - Test Suite ------------------------------------------------------------------------------------
 */

func TestNewCpuMon(t *testing.T) {
	runner := TestRunner{0}
	parser := parse.TestParser{}
	cpu := NewCpuMon(runner, parser)

	if &cpu == nil {
		t.Fail()
	}
}

func TestCpuMon_GetCpus(t *testing.T) {
	runner := TestRunner{0}
	parser := parse.TestParser{}
	cpuMon:= NewCpuMon(runner, parser)

	cpus, err := cpuMon.GetCpus()

	assert.Nil(t, err)
	assert.NotNil(t, cpus)
	assert.Equal(t, 1, len(cpus))

	for i := range cpus {
		assert.Equal(t, "cpu" + strconv.Itoa(i), cpus[i].Name)
	}

	assert.Equal(t, float64(0.7), cpus[0].Usage)
}