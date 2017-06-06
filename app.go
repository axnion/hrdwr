package main

import (
	"github.com/axnion/hrdwr/units"
	"github.com/axnion/hrdwr/util"
	"fmt"
)

func main() {
	runner := util.RealRunner{}
	cpu := new(units.Cpu)
	cpu.SetRunner(runner)
	cpus, _ := cpu.GetCpus()
	fmt.Println(string(cpus))
}