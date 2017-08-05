package main

import (
	"fmt"
	"time"

	"github.com/axnion/hrdwr"
)

func main() {
	for {
		printCPU()
		printDisks()
		printMemory()
		printSensors()
		fmt.Println("\n/////////////////////////////////////////////////////////////////////\n")

		time.Sleep(time.Second)
	}
}

func printCPU() {
	fmt.Println("CPU------------------------------------------------------------------------------")
	cpus, err := hrdwr.GetCpus()

	if err != nil {
		return
	}

	for _, cpu := range cpus {
		fmt.Printf("%s - %.0f\n", cpu.Name, (cpu.Usage * 100))
	}
}

func printDisks() {
	fmt.Print("\n")
	fmt.Println("Disks----------------------------------------------------------------------------")

	disks, err := hrdwr.GetDisks()

	if err != nil {
		return
	}

	for _, disk := range disks {
		fmt.Printf("%s: %v / %v\n", disk.Name, disk.Total, disk.Used)
	}
}

func printMemory() {
	fmt.Print("\n")
	fmt.Println("Memory---------------------------------------------------------------------------")
	mem, err := hrdwr.GetMemory()

	if err != nil {
		return
	}

	fmt.Printf("%v / %v\n", mem.Used, mem.Total)
}

func printSensors() {
	fmt.Print("\n")
	fmt.Println("Sensors--------------------------------------------------------------------------")
	sensors := hrdwr.GetSensors()

	fmt.Println("Temperature")
	for _, temp := range sensors.Temps {
		fmt.Printf("%s: %v\n", temp.Label, temp.Value)
	}
	fmt.Print("\n")

	fmt.Println("Fans")
	for _, fan := range sensors.Fans {
		fmt.Printf("%s: %v\n", fan.Label, fan.Value)
	}
	fmt.Print("\n")

	fmt.Println("Voltage")
	for _, volt := range sensors.Volt {
		fmt.Printf("%s: %v\n", volt.Label, volt.Value)
	}
}
