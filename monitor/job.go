package monitor

import (
	"log"
	"net/http"
	"time"

	"github.com/maxcnunes/go-uptime-api/monitor/data"
)

// Job ...
type Job struct {
	data         *data.DataMonitor
	checkAtEvery time.Duration
}

func (j Job) checkTargetsStatus() {
	results := AsyncHTTPGets(j.data.Target.GetAllURLS())
	for _, result := range results {
		status := http.StatusBadGateway
		if result.Response != nil {
			log.Printf("%s status: %s\n", result.URL, result.Response.Status)
			status = result.Response.StatusCode
		}

		j.saveTracking(result.URL, status)
	}
}

func (j Job) saveTracking(url string, status int) {
	target := j.data.Target.FindOneByURL(url)
	oldStatus := target.Status
	j.data.Track.Create(*target, status)
	target.Status = status
	j.data.Target.Update(target.ID.Hex(), *target)

	// sends notification only if the status has changed
	if oldStatus != status {
		SendNotificaton(*target)
	}
}

func (j Job) checkTargetsPeriodically() {
	StartEventListener(j.data)

	ticker := time.NewTicker(j.checkAtEvery)
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Printf("Checking %d URLs status...", len(j.data.Target.GetAll()))
				j.checkTargetsStatus()
			}
		}
	}()
}

// Start ...
func (j Job) Start(data *data.DataMonitor, checkAtEvery string) {
	j.data = data

	duration, err := time.ParseDuration(checkAtEvery)
	if err != nil {
		log.Fatalf("Value %v is not a valid duration time", checkAtEvery)
	}
	j.checkAtEvery = duration

	log.Printf("Starting targets checking async (every %s)", checkAtEvery)
	j.checkTargetsPeriodically()
}
