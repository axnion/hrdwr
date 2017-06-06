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

		clear()
		if err != nil {
			log.Fatal(err)
		}

		for _, el := range cpus {
			fmt.Printf("%s: %d \n", el.Name, int(el.Usage * 100))
		}

		time.Sleep(1 * time.Second)
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}