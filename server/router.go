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
	results := monitor.AsyncHTTPGets(data.GetAllURLS())
	for _, result := range results {
		if result.Response != nil {
			fmt.Printf("%s status: %s\n", result.URL, result.Response.Status)
		}
	}
}

func (r Router) checkTargetsEvery10seconds() {
	// temp examples
	r.data.AddTarget("https://google.com/")
	r.data.AddTarget("http://twitter.com/")

	monitor.StartEventListener(r.data)
	ticker := time.NewTicker(time.Second * 10)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Printf("Checking %d URLs status...", len(r.data.GetAllTargets()))
				checkTargetsStatus(r.data)
			}
		}
	}()
}

func (r Router) listHandler(rw http.ResponseWriter, req *http.Request) {
	j, err := json.Marshal(r.data.GetAllTargets())
	if err != nil {
		panic(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Write(j)
}

func (r Router) createHanler(rw http.ResponseWriter, req *http.Request) {
	var target monitor.Target

	err := json.NewDecoder(req.Body).Decode(&target)
	if err != nil {
		panic(err)
	}

	r.data.AddTarget(target.URL)

	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.WriteHeader(http.StatusCreated)
}

func (r Router) deleteHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	r.data.RemoveTargetByID(vars["id"])

	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.WriteHeader(http.StatusOK)
}

func (r Router) updateHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	var target monitor.Target

	err := json.NewDecoder(req.Body).Decode(&target)
	if err != nil {
		panic(err)
	}

	r.data.UpdateTarget(vars["id"], target)

	if err != nil {
		panic(err)
	}

	log.Printf("Updated target %s with new URL %s", vars["id"], target.URL)

	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.WriteHeader(http.StatusNoContent)
}

// Start ...
func (r Router) Start(data *monitor.DataMonitor) *mux.Router {
	r.data = data

	log.Print("Starting targets checking async (every 10 sec)")
	r.checkTargetsEvery10seconds()

	log.Print("Starting API server")

	router := mux.NewRouter()
	router.HandleFunc("/", r.listHandler).Methods("GET")
	router.HandleFunc("/", r.createHanler).Methods("POST")
	router.HandleFunc("/{id}", r.updateHandler).Methods("PUT")
	router.HandleFunc("/{id}", r.deleteHandler).Methods("DELETE")

	return router
}
