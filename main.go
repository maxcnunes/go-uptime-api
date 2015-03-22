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
	job       = monitor.Job{}
	router    = server.Router{}
	websocket = server.Websocket{}
)

func main() {
	db.Start()
	defer db.Close()

	data.Start(db)

	http.Handle("/", router.Start(&data))
	http.HandleFunc("/ws", websocket.Start(&data))

	job.Start(&data, getTimeTargetsVerification())

	addr := getServiceAddress()
	log.Printf("Server running on http://%s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getTimeTargetsVerification() string {
	if env := os.Getenv("CHECK_TARGETS_AT_EVERY"); env != "" {
		return env
	}

	return "10m"
}

func getServiceAddress() string {
	if env := os.Getenv("PORT_BEHIND_PROXY"); env != "" {
		return ":" + env
	}
	if env := os.Getenv("VIRTUAL_PORT"); env != "" {
		return ":" + env
	}

	return ":3000"
}
