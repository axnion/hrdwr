package units

import (
	"testing"
)

type TestRunner struct{}

func (TestRunner) Run(cmd string, args ...string) ([]byte, error) {
	return []byte("lol"), nil
}

func TestCpu_GetCpus(t *testing.T) {
	cpu := new(Cpu)
	runner := TestRunner{}
	out, err := cpu.GetCpus(runner)

	if err != nil {
		t.Fail()
	}

	if string(out) != "lol" {
		t.Fail()
	}
}