package main

import (
	"flag"
	"github.com/maxcnunes/monitor-api/monitor"
	"github.com/maxcnunes/monitor-api/server"
	"log"
	"net/http"
)

var (
	db        = monitor.DB{}
	data      = monitor.DataMonitor{}
	router    = server.Router{}
	websocket = server.Websocket{}
	addr      = flag.String("addr", ":3000", "http service address")
)

func main() {
	flag.Parse()

	db.Start()
	defer db.Close()

	data.Start(db)

	http.Handle("/", router.Start(&data))
	http.HandleFunc("/ws", websocket.Start(&data))

	log.Printf("Server running on http://0.0.0.0:%d", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
