package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/maxcnunes/monitor-api/monitor"
	"log"
	"net/http"
)

// Router ...
type Router struct {
	data *monitor.DataMonitor
}

func (r Router) listHandler(rw http.ResponseWriter, req *http.Request) {
	j, err := json.Marshal(r.data.GetAllTargets())
	if err != nil {
		panic(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(j)
}

func (r Router) createHanler(rw http.ResponseWriter, req *http.Request) {
	var target monitor.Target

	err := json.NewDecoder(req.Body).Decode(&target)
	if err != nil {
		panic(err)
	}

	r.data.AddTarget(target.URL)

	rw.WriteHeader(http.StatusCreated)
}

func (r Router) deleteHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	r.data.RemoveTargetByID(vars["id"])

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

	rw.WriteHeader(http.StatusNoContent)
}

// Start ...
func (r Router) Start(data *monitor.DataMonitor) *mux.Router {
	r.data = data

	log.Print("Starting API server")

	router := mux.NewRouter()
	router.HandleFunc("/", addDefaultHeaders(r.listHandler)).Methods("GET", "OPTIONS")
	router.HandleFunc("/", addDefaultHeaders(r.createHanler)).Methods("POST")
	router.HandleFunc("/{id}", addDefaultHeaders(r.updateHandler)).Methods("PUT")
	router.HandleFunc("/{id}", addDefaultHeaders(r.deleteHandler)).Methods("DELETE")

	return router
}

func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		fn(w, r)
	}
}
