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
		printCpus(cpuMon.GetCpus()) 	// CPU
		printMemory(memMon.GetMemory())	// Memory
		printDisk(diskMon.GetDisks())	// Disk

		time.Sleep(1 * time.Second)
	}
}

func printCpus(cpus []libs.CPU, err error) {
	clear()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	fmt.Println("CPU--------------------------------------------------------------------------")

	for _, el := range cpus {
		fmt.Printf("%s: %d \n", el.Name, int(el.Usage * 100))
	}
}

func printMemory(mem libs.Memory, err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	fmt.Println("\nMEMORY-----------------------------------------------------------------------")
	fmt.Printf("Total: %d\n", mem.Total)
	fmt.Printf("Available: %d\n", mem.Available)
}

func printDisk(disks []libs.Disk, err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	fmt.Println("\nDISK-------------------------------------------------------------------------")
	for _, disk := range disks {
		fmt.Printf("%s: \tTotal:%d \tUsed: %d\n", disk.Name, disk.Total, disk.Used)
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}