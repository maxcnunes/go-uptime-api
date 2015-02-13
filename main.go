package main

import (
	"fmt"
	"github.com/maxcnunes/monitor/monitor"
	"time"
)

var (
	data monitor.DataMonitor
)

func checkTargesStatus(data *monitor.DataMonitor) {
	results := monitor.AsyncHTTPGets(data.URLS)
	for _, result := range results {
		if result.Response != nil {
			fmt.Printf("%s status: %s\n", result.URL, result.Response.Status)
		}
	}
}

func main() {
	data = monitor.DataMonitor{}
	// temp examples
	data.URLS = append(data.URLS, "https://google.com/", "http://twitter.com/")

	monitor.StartEventListener(&data)
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			fmt.Printf("Checking %d URLs status...", len(data.URLS))
			checkTargesStatus(&data)
		}
	}
}
