package main

import (
	"flag"
	"net/http"
	"github.com/gorilla/websocket"
	"log"
	"time"
	"github.com/axnion/hrdwr/lib"
)

type payload struct {
	Cpus 	[]lib.CPU
	Disks 	[]lib.Disk
	Memory 	lib.Memory
	Sensors lib.Sensors
}

var addr = flag.String("addr", ":8080", "http service address")
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/", serveClient)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWebSocket(w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func serveClient(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	http.ServeFile(w, r, "client/index.html")
}

func serveWebSocket(w http.ResponseWriter, r *http.Request) {
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