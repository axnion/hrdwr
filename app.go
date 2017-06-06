package main

import (
	"github.com/axnion/hrdwr/units"
	"github.com/axnion/hrdwr/util"
)

func main() {
	runner := util.RealRunner{}
	cpu := new(units.CpuMon)
	cpu.SetRunner(runner)

	cpu.GetCpus()
	//fmt.Println(string(cpus))
}