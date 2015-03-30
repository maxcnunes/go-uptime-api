package monitor

import (
	"log"
	"time"

	"github.com/maxcnunes/monitor-api/monitor/data"
)

// Job ...
type Job struct {
	data         *data.DataMonitor
	checkAtEvery time.Duration
}

func (j Job) checkTargetsStatus() {
	results := AsyncHTTPGets(j.data.Target.GetAllURLS())
	for _, result := range results {
		if result.Response != nil {
			log.Printf("%s status: %s\n", result.URL, result.Response.Status)
		}
	}
}

func (j Job) checkTargetsPeriodically() {
	// temp examples
	j.data.Target.Create("https://google.com/")
	j.data.Target.Create("http://twitter.com/")

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
