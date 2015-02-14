package main

import (
	"github.com/maxcnunes/monitor/server"
)

var (
	router server.Router
)

func main() {
	router = server.Router{}
	router.Start()
}
