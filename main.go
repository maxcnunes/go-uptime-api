package main

import (
	"github.com/maxcnunes/monitor-api/server"
)

var (
	router server.Router
)

func main() {
	router = server.Router{}
	router.Start()
}
