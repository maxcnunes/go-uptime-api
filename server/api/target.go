package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maxcnunes/monitor-api/monitor"
)

// TargetAPI ...
type TargetAPI struct {
	data *monitor.DataMonitor
}

// Start ...
func (api *TargetAPI) Start(data *monitor.DataMonitor) {
	api.data = data
}

// ListHandler ...
func (api *TargetAPI) ListHandler(rw http.ResponseWriter, req *http.Request) {
	j, err := json.Marshal(api.data.GetAllTargets())
	if err != nil {
		panic(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(j)
}

// CreateHanler ...
func (api *TargetAPI) CreateHanler(rw http.ResponseWriter, req *http.Request) {
	var target monitor.Target

	err := json.NewDecoder(req.Body).Decode(&target)
	if err != nil {
		panic(err)
	}

	api.data.AddTarget(target.URL)

	rw.WriteHeader(http.StatusCreated)
}

// DeleteHandler ...
func (api *TargetAPI) DeleteHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	api.data.RemoveTargetByID(vars["id"])

	rw.WriteHeader(http.StatusOK)
}

// UpdateHandler ...
func (api *TargetAPI) UpdateHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	var target monitor.Target

	err := json.NewDecoder(req.Body).Decode(&target)
	if err != nil {
		panic(err)
	}

	api.data.UpdateTarget(vars["id"], target)

	if err != nil {
		panic(err)
	}

	log.Printf("Updated target %s with new URL %s", vars["id"], target.URL)

	rw.WriteHeader(http.StatusOK)
}
