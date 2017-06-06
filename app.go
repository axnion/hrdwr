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
	cpuMon := libs.NewCpuMon(runner)
	memMon := libs.NewMemMon(runner)
	diskMon := libs.NewDiskMon(runner)

	for true {
		// CPU
		cpus, err := cpuMon.GetCpus()
		clear()

		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		fmt.Println("CPU")
		for _, el := range cpus {
			fmt.Printf("%s: %d \n", el.Name, int(el.Usage * 100))
		}

		// MEMORY
		mem, err := memMon.GetMemory()

		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		fmt.Println("\nMEMORY")
		fmt.Printf("Total: %d\n", mem.Total)
		fmt.Printf("Available: %d\n", mem.Available)

		// DISK
		disks, err := diskMon.GetDisks()

		fmt.Println("\nDISK")

		for _, disk := range disks {
			fmt.Printf("%s: Total: %d Used: %d", disk.Name, disk.Total, disk.Used)
		}

		time.Sleep(1 * time.Second)
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}