package units

import (
	"testing"
)

type TestRunner struct{}

func (TestRunner) Run(cmd string, args ...string) ([]byte, error) {
	return []byte("cpuinfo file content"), nil
}

func TestCpu_GetCpus(t *testing.T) {
	runner := TestRunner{}
	cpu := new(Cpu)
	cpu.SetRunner(runner)
	cores, err := cpu.GetCpus()

	if err != nil {
		t.Fail()
	}

	if string(cores) != "cpuinfo file content" {
		t.Fail()
	}
}

func TestCpu_getCpuinfo(t *testing.T) {
	runner := TestRunner{}
	out, err := getCpuinfo(runner)

	if err != nil {
		t.Fail()
	}

	if string(out) != "cpuinfo file content" {
		t.Fail()
	}
}