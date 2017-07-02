package server

import (
	"log"
	"github.com/gorilla/websocket"
	"github.com/axnion/hrdwr/server/lib"
	"time"
	"net/http"
)

type payload struct {
	Cpus 	[]lib.CPU
	Disks 	[]lib.Disk
	Memory 	lib.Memory
	Sensors lib.Sensors
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	go streamStats(conn)
}

func streamStats(conn *websocket.Conn) {
	for {
		cpus, err := lib.GetCpus()
		if err != nil {log.Fatal("Error fetching CPU data")}

		disks, err := lib.GetDisks()
		if err != nil {log.Fatal("Error fetching disk data")}

		memory, err := lib.GetMemory()
		if err != nil {log.Fatal("Error fetching memory data")}

		sensors := lib.GetSensors()

		data := payload {
			cpus,
			disks,
			memory,
			sensors,
		}



		conn.WriteJSON(data)
		time.Sleep(time.Second)
	}
}
