package main

import (
	"flag"
	"github.com/maxcnunes/monitor-api/monitor"
	"github.com/maxcnunes/monitor-api/server"
	"log"
	"net/http"
)

var (
	data      = monitor.DataMonitor{}
	router    = server.Router{}
	websocket = server.Websocket{}
	addr      = flag.String("addr", ":3000", "http service address")
)

func main() {
	flag.Parse()

	data.Events = make(chan int)

	http.Handle("/", router.Start(&data))
	http.HandleFunc("/ws", websocket.Start(&data))

	log.Printf("Server running on http://0.0.0.0:%d", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
