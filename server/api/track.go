package api

import (
	"encoding/json"
	"net/http"

	"github.com/maxcnunes/go-uptime-api/monitor/data"
)

// TrackAPI ...
type TrackAPI struct {
	data *data.DataMonitor
}

// Start ...
func (api *TrackAPI) Start(data *data.DataMonitor) {
	api.data = data
}

// ListHandler ...
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
