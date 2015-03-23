package server

import (
	"log"
	"net/http"

	"github.com/maxcnunes/monitor-api/server/api"

	"github.com/gorilla/mux"
	"github.com/maxcnunes/monitor-api/monitor"
)

// Router ...
type Router struct{}

var (
	targetAPI = api.TargetAPI{}
)

// Start ...
func (r *Router) Start(data *monitor.DataMonitor) *mux.Router {
	log.Print("Starting API server")
	router := mux.NewRouter()

	// targets api
	targetAPI.Start(data)
	router.HandleFunc("/targets", addDefaultHeaders(targetAPI.ListHandler)).Methods("GET", "OPTIONS")
	router.HandleFunc("/targets", addDefaultHeaders(targetAPI.CreateHanler)).Methods("POST")
	router.HandleFunc("/targets/{id}", addDefaultHeaders(targetAPI.UpdateHandler)).Methods("PUT")
	router.HandleFunc("/targets/{id}", addDefaultHeaders(targetAPI.DeleteHandler)).Methods("DELETE")

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
