package main

import (
	"log"

	"github.com/influxdata/influxdb/client/v2"
)

const addr = "http://influxdb:8086"
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

	batch, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  db,
		Precision: "s",
	})

	err = influx.Write(batch)

	if err != nil {
		log.Fatal(err)
	}
}
