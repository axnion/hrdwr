package server

import (
	"github.com/axnion/hrdwr/lib"
	"fmt"
	"log"
	"os/exec"
	"os"
	"time"
)

func main() {
	for true {
		printCpus(lib.GetCpus())             	// CPU
		printMemory(lib.GetMemory())      	// Memory
		printDisk(lib.GetDisks())        	// Disk
		printSensors(lib.GetSensors()) 		// Sensors (temp, fans, volt)

		time.Sleep(1 * time.Second)		// Wait 1 second
	}
}

func printCpus(cpus []lib.CPU, err error) {
	clear()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	fmt.Println("CPU--------------------------------------------------------------------------")

	for _, el := range cpus {
		fmt.Printf("%s: %d\n", el.Name, int(el.Usage * 100))
	}
}

func printMemory(mem lib.Memory, err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	fmt.Println("\nMEMORY-----------------------------------------------------------------------")
	fmt.Printf("Total: %d\n", mem.Total)
	fmt.Printf("Used: %d\n", mem.Used)
}

func printDisk(disks []lib.Disk, err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	fmt.Println("\nDISK-------------------------------------------------------------------------")
	for _, disk := range disks {
		fmt.Printf("%s: \tTotal:%d \tUsed: %d\n", disk.Name, disk.Total, disk.Used)
	}
}

func printSensors(sensors lib.Sensors) {
	fmt.Println("\nSensors-----------------------------------------------------------------------")
	fmt.Print("Temperatures")
	for _, temp := range sensors.Temps {
		fmt.Printf("\n\t%v: %v", temp.Label, temp.Value)
	}

	fmt.Print("\nFans")
	for _, fans := range sensors.Fans{
		fmt.Printf("\n\t%v: %v", fans.Label, fans.Value)
	}

	fmt.Print("\nVoltage")
	for _, volt:= range sensors.Volt {
		fmt.Printf("\n\t%v: %v", volt.Label, volt.Value)
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}