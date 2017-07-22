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

		cpus, err := lib.GetCpus()

		if err != nil {
			log.Fatal(err)
			return
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

			batch.AddPoint(point)
		}

		err = influx.Write(batch)

		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second)
	}
}
