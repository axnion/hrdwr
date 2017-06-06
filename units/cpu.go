package units

import "github.com/axnion/hrdwr/util"

type Cpu struct{}

func (cpu Cpu) GetCpus(runner util.Runner) ([]byte, error) {
	return runner.Run("echo", "Hello World")
}