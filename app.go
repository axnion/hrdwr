package main

import (
	"flag"
	"net/http"
	"log"
	"github.com/axnion/hrdwr/server"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	http.HandleFunc("/", serveClient)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server.ServeWebSocket(w, r)
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
