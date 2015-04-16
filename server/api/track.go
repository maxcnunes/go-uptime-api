package api

import (
	"encoding/json"
	"net/http"

	"github.com/maxcnunes/go-uptime-api/monitor/data"
)

// TrackAPI has all routes related to track data manipulation
type TrackAPI struct {
	data *data.DataMonitor
}

// Start a new instance of track api
func (api *TrackAPI) Start(data *data.DataMonitor) {
	api.data = data
}

// ListHandler handles GET request returning all tracks
func (api *TrackAPI) ListHandler(rw http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	targetID := ""
	if len(query["targetId"]) > 0 {
		targetID = query["targetId"][0]
	}

	j, err := json.Marshal(api.data.Track.Find(targetID))
	if err != nil {
		panic(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(j)
}
