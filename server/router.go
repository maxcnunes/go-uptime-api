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

// Router ...
type Router struct {
	data *monitor.DataMonitor
}

func checkTargetsStatus(data *monitor.DataMonitor) {
	results := monitor.AsyncHTTPGets(data.URLS)
	for _, result := range results {
		if result.Response != nil {
			fmt.Printf("%s status: %s\n", result.URL, result.Response.Status)
		}
	}
}

func (r Router) checkTargetsEvery10seconds() {
	// temp examples
	r.data.URLS = append(r.data.URLS, "https://google.com/", "http://twitter.com/")

	monitor.StartEventListener(r.data)
	ticker := time.NewTicker(time.Second * 10)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Printf("Checking %d URLs status...", len(r.data.URLS))
				checkTargetsStatus(r.data)
			}
		}
	}()
}

func (r Router) listHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	j, err := json.Marshal(r.data.URLS)
	if err != nil {
		panic(err)
	}

	rw.Write(j)
}

// Start ...
func (r Router) Start(data *monitor.DataMonitor) *mux.Router {
	r.data = data

	log.Print("Starting targets checking async (every 10 sec)")
	r.checkTargetsEvery10seconds()

	log.Print("Starting API server")

	router := mux.NewRouter()
	router.HandleFunc("/", r.listHandler).Methods("GET")

	return router
}
