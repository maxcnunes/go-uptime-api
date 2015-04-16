package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maxcnunes/go-uptime-api/monitor/data"
	"github.com/maxcnunes/go-uptime-api/monitor/entities"
)

// TargetAPI has all routes related to target data manipulation
type TargetAPI struct {
	data *data.DataMonitor
}

// Start a new instance of target api
func (api *TargetAPI) Start(data *data.DataMonitor) {
	api.data = data
}

// ListHandler handles GET request returning all targets
func (api *TargetAPI) ListHandler(rw http.ResponseWriter, req *http.Request) {
	j, err := json.Marshal(api.data.Target.GetAll())
	if err != nil {
		panic(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(j)
}

// DetailHandler handles GET request returning a single target
func (api *TargetAPI) DetailHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	j, err := json.Marshal(api.data.Target.FindOneByID(vars["id"]))
	if err != nil {
		panic(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(j)
}

// CreateHanler handles POST request creating a new target
func (api *TargetAPI) CreateHanler(rw http.ResponseWriter, req *http.Request) {
	var target entities.Target

	err := json.NewDecoder(req.Body).Decode(&target)
	if err != nil {
		panic(err)
	}

	api.data.Target.Create(target.URL, target.Emails)

	rw.WriteHeader(http.StatusCreated)
}

// DeleteHandler handles DELETE request deleting the proper target
func (api *TargetAPI) DeleteHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	targetID := vars["id"]
	api.data.Target.RemoveByID(targetID)
	api.data.Track.RemoveByTargetID(targetID)

	rw.WriteHeader(http.StatusOK)
}

// UpdateHandler handles PUT request updating the proper target
func (api *TargetAPI) UpdateHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	var target entities.Target

	err := json.NewDecoder(req.Body).Decode(&target)
	if err != nil {
		panic(err)
	}

	api.data.Target.Update(vars["id"], target)

	if err != nil {
		panic(err)
	}

	log.Printf("Updated target %s with new URL %s", vars["id"], target.URL)

	rw.WriteHeader(http.StatusOK)
}
