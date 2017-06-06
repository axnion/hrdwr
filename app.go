package main

import (
	"github.com/axnion/hrdwr/units"
	"github.com/axnion/hrdwr/util"
	"fmt"
)

func main() {
	runner := util.RealRunner{}
	cpu := new(units.CpuMon)
	cpu.SetRunner(runner)
	cpus := []units.CPU{}
	cpus, _ = cpu.GetCpus(cpus)

	for _, el := range cpus {
		fmt.Printf("%s: %f\n", el.Name, el.Usage)
	}
}