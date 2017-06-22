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

func (runner CpuTestRunner) Run(cmd string, args ...string) ([]byte, error) {
	runner.executions++

	if(runner.executions % 2 == 0) {
		return []byte(runner.stat1), runner.err
	} else {
		return []byte(runner.stat2), runner.err
	}
}

/*
 * - Test Suite ------------------------------------------------------------------------------------
 */

func prepareRunner(stat1 []byte, stat2 []byte, err error) {
	runner = CpuTestRunner{0, stat1, stat2, err}
}

func TestGetCpus(t *testing.T) {
	stat1 := "cpu  763492 664 124647 6661572 2929 0 8669 0 0 0\n" +
		"cpu0 113887 105 13672 794718 1125 0 5695 0 0 0\n" +
		"cpu1 90391 71 19293 836664 354 0 1653 0 0 0\n"

	stat2 := "cpu  763517 664 124651 6661942 2929 0 8670 0 0 0\n" +
		"cpu0 113891 105 13673 794762 1125 0 5696 0 0 0\n" +
		"cpu1 90392 71 19294 836711 354 0 1653 0 0 0\n"

	prepareRunner([]byte(stat1), []byte(stat2), nil)

	cpus, err := GetCpus()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 2, len(cpus))

	assert.Equal(t, "cpu0", cpus[0].Name)
	assert.Equal(t, "cpu1", cpus[1].Name)

	assert.Equal(t, float64(10), cpus[0].Usage)
}
