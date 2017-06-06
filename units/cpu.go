package units

import "github.com/axnion/hrdwr/util"

type Cpu struct{
	cpuinfo []byte
}

type Core struct {
	id string
}

var runner util.Runner

func (cpu Cpu) GetCpus() ([]byte, error) {
	if cpu.cpuinfo == nil {
		cpu.cpuinfo, _ = getCpuinfo(runner)
	}

	return cpu.cpuinfo, nil
}

func (cpu Cpu) SetRunner(newRunner util.Runner) {
	runner = newRunner
}

func parseCpuinfo(content []byte) ([]Core) {

}

func getCpuinfo(runner util.Runner) ([]byte, error) {
	return runner.Run("cat", "/proc/cpuinfo")
}