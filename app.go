package main

import (
	"flag"
	"net/http"
	"log"
	"github.com/axnion/hrdwr/server"
)

var addr = flag.String("addr", ":8080", "http service address")
var clientBuild = "client/build/"

func main() {
	fs := http.FileServer(http.Dir(clientBuild + "static"))

	http.HandleFunc("/", serveClient)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server.ServeWebSocket(w, r)
	})

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func serveClient(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	http.ServeFile(w, r, clientBuild + "index.html")
}
