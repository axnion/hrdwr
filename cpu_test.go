package hrdwr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * - Test Runner -----------------------------------------------------------------------------------
 */

/**
 * A struct which replaces the original Runner object used in the CPU test package.
 */
type CpuTestRunner struct {
	executions int
	stat1      []byte
	stat2      []byte
	err1       error
	err2       error
}

/**
 * The mocked run method of the CpuTestRunner struct. Counts the executions and returns two different byte arrays
 * depending on how many executions have been made, this is to simulate two readings of the /proc/stat file.
 */
func (runner *CpuTestRunner) run(cmd string, args ...string) ([]byte, error) {
	runner.executions++

	if runner.executions%2 == 0 {
		return []byte(runner.stat2), runner.err1
	} else {
		return []byte(runner.stat1), runner.err2
	}

}

/*
 * - Test Suite ------------------------------------------------------------------------------------
 */

/**
 * Sets the global runner variable to use a CpuTestRunner pointer and also set the data to be used by CpuTestRunner.
 */
func prepareCPURunner(stat1 []byte, stat2 []byte, err1 error, err2 error) {
	runner = &CpuTestRunner{0, stat1, stat2, err1, err2}
}

/**
 * The mail test case for GetCpus function. Two example readings from an /proc/stat file is sent to the CpuTestRunner
 * to be returned later. Then the function is called and the results are checked so they match what is expected.
 */
func TestGetCpus(t *testing.T) {
	stat1 := "cpu  207034 80 33423 1623348 1809 0 1643 0 0 0\n" +
		"cpu0 27995 10 3813 194891 517 0 856 0 0 0\n" +
		"cpu1 27219 9 3837 202513 285 0 478 0 0 0\n"

	stat2 := "cpu  207180 80 33435 1623603 1809 0 1644 0 0 0\n" +
		"cpu0 28017 10 3814 194916 517 0 856 0 0 0\n" +
		"cpu1 27232 9 3839 202548 285 0 478 0 0 0\n"

	prepareCPURunner([]byte(stat1), []byte(stat2), nil, nil)

	cpus, err := GetCpus()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(cpus))
	assert.Equal(t, "cpu0", cpus[0].Name)
	assert.Equal(t, "cpu1", cpus[1].Name)
	assert.Equal(t, float64(0.4791666666666667), cpus[0].Usage)
	assert.Equal(t, float64(0.3), cpus[1].Usage)
}

/**
 * Tests if runner experiences an error during the first execution.
 */
func TestGetCpusFirstRunError(t *testing.T) {
	stat1 := "cpu  207034 80 33423 1623348 1809 0 1643 0 0 0\n" +
		"cpu0 27995 10 3813 194891 517 0 856 0 0 0\n" +
		"cpu1 27219 9 3837 202513 285 0 478 0 0 0\n"

	stat2 := "cpu  207180 80 33435 1623603 1809 0 1644 0 0 0\n" +
		"cpu0 28017 10 3814 194916 517 0 856 0 0 0\n" +
		"cpu1 27232 9 3839 202548 285 0 478 0 0 0\n"

	err := errors.New("test error 1")
	prepareCPURunner([]byte(stat1), []byte(stat2), err, nil)

	cpus, err := GetCpus()

	assert.Nil(t, cpus)
	assert.NotNil(t, err)
	assert.Equal(t, "test error 1", err.Error())
}

/**
 * Tests if runner experiences an error during the second execution.
 */
func TestGetCpusSecondRunError(t *testing.T) {
	stat1 := "cpu  207034 80 33423 1623348 1809 0 1643 0 0 0\n" +
		"cpu0 27995 10 3813 194891 517 0 856 0 0 0\n" +
		"cpu1 27219 9 3837 202513 285 0 478 0 0 0\n"

	stat2 := "cpu  207180 80 33435 1623603 1809 0 1644 0 0 0\n" +
		"cpu0 28017 10 3814 194916 517 0 856 0 0 0\n" +
		"cpu1 27232 9 3839 202548 285 0 478 0 0 0\n"

	err := errors.New("test error 2")
	prepareCPURunner([]byte(stat1), []byte(stat2), nil, err)

	cpus, err := GetCpus()

	assert.Nil(t, cpus)
	assert.NotNil(t, err)
	assert.Equal(t, "test error 2", err.Error())
}

/**
 * Tests if runner experiences an error on the user column during the first parsing.
 */
func TestGetCpusFirstParsingUserError(t *testing.T) {
	stat1 := "cpu  207034 80 33423 1623348 1809 0 1643 0 0 0\n" +
		"cpu0 27aa995 10 3813 194891 517 0 856 0 0 0\n" +
		"cpu1 27219 9 3837 202513 285 0 478 0 0 0\n"

	stat2 := "cpu  207180 80 33435 1623603 1809 0 1644 0 0 0\n" +
		"cpu0 28017 10 3814 194916 517 0 856 0 0 0\n" +
		"cpu1 27232 9 3839 202548 285 0 478 0 0 0\n"

	prepareCPURunner([]byte(stat1), []byte(stat2), nil, nil)

	cpus, err := GetCpus()

	assert.Nil(t, cpus)
	assert.NotNil(t, err)
}

/**
 * Tests if runner experiences an error on the user column during the second parsing.
 */
func TestGetCpusSecondParsingUserError(t *testing.T) {
	stat1 := "cpu  207034 80 33423 1623348 1809 0 1643 0 0 0\n" +
		"cpu0 27995 10 3813 194891 517 0 856 0 0 0\n" +
		"cpu1 27219 9 3837 202513 285 0 478 0 0 0\n"

	stat2 := "cpu  207180 80 33435 1623603 1809 0 1644 0 0 0\n" +
		"cpu0 28a017 10 3814 194916 517 0 856 0 0 0\n" +
		"cpu1 27232 9 3839 202548 285 0 478 0 0 0\n"

	prepareCPURunner([]byte(stat1), []byte(stat2), nil, nil)

	cpus, err := GetCpus()

	assert.Nil(t, cpus)
	assert.NotNil(t, err)
}

/**
 * Tests if runner experiences an error on the nice column
 */
func TestGetCpusParsingNiceError(t *testing.T) {
	stat1 := "cpu  207034 80 33423 1623348 1809 0 1643 0 0 0\n" +
		"cpu0 27995 1i0 3813 194891 517 0 856 0 0 0\n" +
		"cpu1 27219 9 3837 202513 285 0 478 0 0 0\n"

	stat2 := "cpu  207180 80 33435 1623603 1809 0 1644 0 0 0\n" +
		"cpu0 28017 10 3814 194916 517 0 856 0 0 0\n" +
		"cpu1 27232 9 3839 202548 285 0 478 0 0 0\n"

	prepareCPURunner([]byte(stat1), []byte(stat2), nil, nil)

	cpus, err := GetCpus()

	assert.Nil(t, cpus)
	assert.NotNil(t, err)
}

/**
 * Tests if runner experiences an error on the system column
 */
func TestGetCpusParsingSystemError(t *testing.T) {
	stat1 := "cpu  207034 80 33423 1623348 1809 0 1643 0 0 0\n" +
		"cpu0 27995 10 38i13 194891 517 0 856 0 0 0\n" +
		"cpu1 27219 9 3837 202513 285 0 478 0 0 0\n"

	stat2 := "cpu  207180 80 33435 1623603 1809 0 1644 0 0 0\n" +
		"cpu0 28017 10 3814 194916 517 0 856 0 0 0\n" +
		"cpu1 27232 9 3839 202548 285 0 478 0 0 0\n"

	prepareCPURunner([]byte(stat1), []byte(stat2), nil, nil)

	cpus, err := GetCpus()

	assert.Nil(t, cpus)
	assert.NotNil(t, err)
}

/**
 * Tests if runner experiences an error on the idle column
 */
func TestGetCpusParsingIdleError(t *testing.T) {
	stat1 := "cpu  207034 80 33423 1623348 1809 0 1643 0 0 0\n" +
		"cpu0 27995 10 3813 1948i91 517 0 856 0 0 0\n" +
		"cpu1 27219 9 3837 202513 285 0 478 0 0 0\n"

	stat2 := "cpu  207180 80 33435 1623603 1809 0 1644 0 0 0\n" +
		"cpu0 28017 10 3814 194916 517 0 856 0 0 0\n" +
		"cpu1 27232 9 3839 202548 285 0 478 0 0 0\n"

	prepareCPURunner([]byte(stat1), []byte(stat2), nil, nil)

	cpus, err := GetCpus()

	assert.Nil(t, cpus)
	assert.NotNil(t, err)
}

/**
 * Tests if runner experiences an error on the iowait column
 */
func TestGetCpusParsingIowaitError(t *testing.T) {
	stat1 := "cpu  207034 80 33423 1623348 1809 0 1643 0 0 0\n" +
		"cpu0 27995 10 3813 194891 51i7 0 856 0 0 0\n" +
		"cpu1 27219 9 3837 202513 285 0 478 0 0 0\n"

	stat2 := "cpu  207180 80 33435 1623603 1809 0 1644 0 0 0\n" +
		"cpu0 28017 10 3814 194916 517 0 856 0 0 0\n" +
		"cpu1 27232 9 3839 202548 285 0 478 0 0 0\n"

	prepareCPURunner([]byte(stat1), []byte(stat2), nil, nil)

	cpus, err := GetCpus()

	assert.Nil(t, cpus)
	assert.NotNil(t, err)
}

/**
 * Tests if runner experiences an error on the irq column
 */
func TestGetCpusParsingIrqError(t *testing.T) {
	stat1 := "cpu  207034 80 33423 1623348 1809 0 1643 0 0 0\n" +
		"cpu0 27995 10 3813 194891 517 0i 856 0 0 0\n" +
		"cpu1 27219 9 3837 202513 285 0 478 0 0 0\n"

	stat2 := "cpu  207180 80 33435 1623603 1809 0 1644 0 0 0\n" +
		"cpu0 28017 10 3814 194916 517 0 856 0 0 0\n" +
		"cpu1 27232 9 3839 202548 285 0 478 0 0 0\n"

	prepareCPURunner([]byte(stat1), []byte(stat2), nil, nil)

	cpus, err := GetCpus()

	assert.Nil(t, cpus)
	assert.NotNil(t, err)
}

/**
 * Tests if runner experiences an error on the softirq column
 */
func TestGetCpusParsingSoftirqError(t *testing.T) {
	stat1 := "cpu  207034 80 33423 1623348 1809 0 1643 0 0 0\n" +
		"cpu0 27995 10 3813 194891 517 0 85i6 0 0 0\n" +
		"cpu1 27219 9 3837 202513 285 0 478 0 0 0\n"

	stat2 := "cpu  207180 80 33435 1623603 1809 0 1644 0 0 0\n" +
		"cpu0 28017 10 3814 194916 517 0 856 0 0 0\n" +
		"cpu1 27232 9 3839 202548 285 0 478 0 0 0\n"

	prepareCPURunner([]byte(stat1), []byte(stat2), nil, nil)

	cpus, err := GetCpus()

	assert.Nil(t, cpus)
	assert.NotNil(t, err)
}

/**
 * Tests if runner experiences an error on the steal column
 */
func TestGetCpusParsingStealError(t *testing.T) {
	stat1 := "cpu  207034 80 33423 1623348 1809 0 1643 0 0 0\n" +
		"cpu0 27995 10 3813 194891 517 0 856 0i 0 0\n" +
		"cpu1 27219 9 3837 202513 285 0 478 0 0 0\n"

	stat2 := "cpu  207180 80 33435 1623603 1809 0 1644 0 0 0\n" +
		"cpu0 28017 10 3814 194916 517 0 856 0 0 0\n" +
		"cpu1 27232 9 3839 202548 285 0 478 0 0 0\n"

	prepareCPURunner([]byte(stat1), []byte(stat2), nil, nil)

	cpus, err := GetCpus()

	assert.Nil(t, cpus)
	assert.NotNil(t, err)
}
