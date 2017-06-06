package main

import (
	"github.com/axnion/hrdwr/libs"
	"fmt"
	"log"
	"time"
	"os/exec"
	"os"
)

func main() {
	runner := libs.RealRunner{}
	cpu := libs.NewCpuMon(runner)

	for true {
		// CPU
		cpus, err := cpu.GetCpus()
		clear()

		if err != nil {
			log.Fatal(err)
		}

		for _, el := range cpus {
			fmt.Printf("%s: %d \n", el.Name, int(el.Usage * 100))
		}

		// MEMORY


		time.Sleep(1 * time.Second)
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}