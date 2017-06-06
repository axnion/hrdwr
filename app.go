package main

import (
	"github.com/axnion/hrdwr/units"
	"github.com/axnion/hrdwr/util"
	"fmt"
)

func main() {
	runner := util.RealRunner{}
	cpu := new(units.Cpu)
	cpus, _ := cpu.GetCpus(runner)
	fmt.Println(string(cpus))
}