package lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

/*
 * - Test Runner -----------------------------------------------------------------------------------
 */

type CpuTestRunner struct {
	executions int
	stat1 []byte
	stat2 []byte
	err error
}

func (runner *CpuTestRunner) run(cmd string, args ...string) ([]byte, error) {
	runner.executions++

	if runner.executions % 2 == 0 {
		return []byte(runner.stat2), runner.err
	} else {
		return []byte(runner.stat1), runner.err
	}

}

/*
 * - Test Suite ------------------------------------------------------------------------------------
 */

func prepareCPURunner(stat1 []byte, stat2 []byte, err error) {
	runner = &CpuTestRunner{0, stat1, stat2, err}
}

func TestGetCpus(t *testing.T) {
	stat1 := "cpu  207034 80 33423 1623348 1809 0 1643 0 0 0\n" +
		"cpu0 27995 10 3813 194891 517 0 856 0 0 0\n" +
		"cpu1 27219 9 3837 202513 285 0 478 0 0 0\n"

	stat2 := "cpu  207180 80 33435 1623603 1809 0 1644 0 0 0\n" +
		"cpu0 28017 10 3814 194916 517 0 856 0 0 0\n" +
		"cpu1 27232 9 3839 202548 285 0 478 0 0 0\n"

	prepareCPURunner([]byte(stat1), []byte(stat2), nil)

	cpus, err := GetCpus()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(cpus))
	assert.Equal(t, "cpu0", cpus[0].Name)
	assert.Equal(t, "cpu1", cpus[1].Name)
	assert.Equal(t, float64(0.4791666666666667), cpus[0].Usage)
	assert.Equal(t, float64(0.3), cpus[1].Usage)
}
