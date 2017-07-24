package main

import (
	"log"
	"time"

	"github.com/axnion/hrdwr/lib"
	"github.com/influxdata/influxdb/client/v2"
)

const addr = "http://localhost:8086"
const db = "hrdwr"

// TODO: Add hostname to data being sent in to influx

func main() {
	influx, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: addr,
	})

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	for {
		batch, err := client.NewBatchPoints(client.BatchPointsConfig{
			Database:  db,
			Precision: "s",
		})

		//diskPoints := getDisks(influx)

		batch.AddPoints(getCpuDataPoints())
		batch.AddPoints(getDiskDataPoints())
		batch.AddPoint(getMemoryDataPoints())
		batch.AddPoints(getSensorDataPoints())

		err = influx.Write(batch)

		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second)
	}

}

func getCpuDataPoints() []*client.Point {
	var points []*client.Point

	cpus, err := lib.GetCpus()

	if err != nil {
		log.Fatal(err)
	}

	for _, cpu := range cpus {
		tags := map[string]string{
			"cpu": cpu.Name,
		}

		fields := map[string]interface{}{
			"usage": cpu.Usage,
		}

		point, err := client.NewPoint("cpu_usage", tags, fields, time.Now())

		if err != nil {
			log.Fatal(err)
		}

		points = append(points, point)
	}

	return points
}

func getDiskDataPoints() []*client.Point {
	var points []*client.Point

	disks, err := lib.GetDisks()

	if err != nil {
		log.Fatal(err)
	}

	for _, disk := range disks {
		tags := map[string]string{
			"disk": disk.Name,
		}

		fields := map[string]interface{}{
			"total": disk.Total,
			"used":  disk.Used,
		}

		point, err := client.NewPoint("disk_utilization", tags, fields, time.Now())

		if err != nil {
			log.Fatal(err)
		}

		points = append(points, point)

	}

	return points
}

func getMemoryDataPoints() *client.Point {
	memory, err := lib.GetMemory()

	if err != nil {
		log.Fatal(err)
	}

	tags := map[string]string{
		"memory": "memory",
	}

	fields := map[string]interface{}{
		"total": memory.Total,
		"used":  memory.Used,
	}

	point, err := client.NewPoint("memory_utlization", tags, fields, time.Now())

	if err != nil {
		log.Fatal(err)
	}

	return point
}

func getSensorDataPoints() []*client.Point {
	var points []*client.Point

	sensors := lib.GetSensors()

	for _, temp := range sensors.Temps {
		tags := map[string]string{
			"sensors_type": "temp",
			"label":        temp.Label,
		}

		fields := map[string]interface{}{
			"value": temp.Value,
		}

		point, err := client.NewPoint("Temperature", tags, fields, time.Now())

		if err != nil {
			log.Fatal(err)
		}

		points = append(points, point)
	}

	for _, fan := range sensors.Fans {
		tags := map[string]string{
			"sensors_type": "fan",
			"label":        fan.Label,
		}

		fields := map[string]interface{}{
			"value": fan.Value,
		}

		point, err := client.NewPoint("Fan Speed", tags, fields, time.Now())

		if err != nil {
			log.Fatal(err)
		}

		points = append(points, point)
	}

	for _, volt := range sensors.Volt {
		tags := map[string]string{
			"sensors_type": "volt",
			"label":        volt.Label,
		}

		fields := map[string]interface{}{
			"value": volt.Value,
		}

		point, err := client.NewPoint("Voltage", tags, fields, time.Now())

		if err != nil {
			log.Fatal(err)
		}

		points = append(points, point)
	}
	return points

}
