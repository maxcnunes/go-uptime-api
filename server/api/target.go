package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maxcnunes/monitor-api/monitor/data"
	"github.com/maxcnunes/monitor-api/monitor/entities"
)

// TargetAPI ...
type TargetAPI struct {
	data *data.DataMonitor
}

// Start ...
func (api *TargetAPI) Start(data *data.DataMonitor) {
	api.data = data
}

// ListHandler ...
func (api *TargetAPI) ListHandler(rw http.ResponseWriter, req *http.Request) {
	j, err := json.Marshal(api.data.Target.GetAll())
	if err != nil {
		panic(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(j)
}

// DetailHandler ...
func (api *TargetAPI) DetailHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	j, err := json.Marshal(api.data.Target.FindOneByID(vars["id"]))
	if err != nil {
		panic(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(j)
}

// CreateHanler ...
func (api *TargetAPI) CreateHanler(rw http.ResponseWriter, req *http.Request) {
	var target entities.Target

	err := json.NewDecoder(req.Body).Decode(&target)
	if err != nil {
		panic(err)
	}

	api.data.Target.Create(target.URL, target.Emails)

	rw.WriteHeader(http.StatusCreated)
}

// DeleteHandler ...
func (api *TargetAPI) DeleteHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	api.data.Target.RemoveByID(vars["id"])

	rw.WriteHeader(http.StatusOK)
}

// UpdateHandler ...
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
