package main

import (
	"github.com/maxcnunes/monitor-api/monitor"
	"github.com/maxcnunes/monitor-api/server"
	"log"
	"net/http"
	"os"
)

var (
	db        = monitor.DB{}
	data      = monitor.DataMonitor{}
	router    = server.Router{}
	websocket = server.Websocket{}
)

func main() {
	var addr string
	if env := os.Getenv("PORT_BEHIND_PROXY"); env != "" {
		addr = ":" + env
	} else if env := os.Getenv("VIRTUAL_PORT"); env != "" {
		addr = ":" + env
	} else {
		addr = ":3000"
	}

	db.Start()
	defer db.Close()

	data.Start(db)

	http.Handle("/", router.Start(&data))
	http.HandleFunc("/ws", websocket.Start(&data))

	log.Printf("Server running on http://%s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
