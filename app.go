package main

import (
	"github.com/axnion/hrdwr/units"
	"github.com/axnion/hrdwr/util"
	"fmt"
	"log"
	"time"
	"os/exec"
	"os"
)

func main() {
	runner := util.RealRunner{}
	cpu := units.NewCpuMon(runner)

	for true {
		cpus, err := cpu.GetCpus()

		if err != nil {
			log.Fatal(err)
		}

		for _, el := range cpus {
			fmt.Printf("%s: %f\n", el.Name, el.Usage)
		}

		time.Sleep(1 * time.Second)
		clear()
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}