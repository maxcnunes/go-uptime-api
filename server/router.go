package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/maxcnunes/monitor-api/monitor"
	"log"
	"net/http"
	"time"
)

var (
	data monitor.DataMonitor
)

// Router ...
type Router struct{}

func checkTargetsStatus(data *monitor.DataMonitor) {
	results := monitor.AsyncHTTPGets(data.URLS)
	for _, result := range results {
		if result.Response != nil {
			fmt.Printf("%s status: %s\n", result.URL, result.Response.Status)
		}
	}
}

func checkTargetsEvery10seconds() {
	data = monitor.DataMonitor{}
	// temp examples
	data.URLS = append(data.URLS, "https://google.com/", "http://twitter.com/")

	monitor.StartEventListener(&data)
	ticker := time.NewTicker(time.Second * 10)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Printf("Checking %d URLs status...", len(data.URLS))
				checkTargetsStatus(&data)
			}
		}
	}()
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	j, err := json.Marshal(data.URLS)
	if err != nil {
		panic(err)
	}

	w.Write(j)
}

// Start ...
func (r Router) Start() {
	log.Print("Starting targets checking async (every 10 sec)")
	checkTargetsEvery10seconds()

	log.Print("Starting server")

	router := mux.NewRouter()
	router.HandleFunc("/", listHandler).Methods("GET")

	http.Handle("/", router)

	log.Print("Server running on http://0.0.0.0:3000")
	http.ListenAndServe(":3000", nil)
}
