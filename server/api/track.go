package api

import (
	"encoding/json"
	"net/http"

	"github.com/maxcnunes/monitor-api/monitor"
)

// TrackAPI ...
type TrackAPI struct {
	data *monitor.DataMonitor
}

// Start ...
func (api *TrackAPI) Start(data *monitor.DataMonitor) {
	api.data = data
}

// ListHandler ...
func (api *TrackAPI) ListHandler(rw http.ResponseWriter, req *http.Request) {
	j, err := json.Marshal(api.data.GetAllTracks())
	if err != nil {
		panic(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(j)
}
