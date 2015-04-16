package data

import "github.com/maxcnunes/go-uptime-api/monitor/entities"

// DataMonitor aggregates the data configuration
// like the database configuration and connections to MongoDB collections
type DataMonitor struct {
	DB     DB
	Events chan entities.Event
	Target *DataTarget
	Track  *DataTrack
}

// Start ...
func (d *DataMonitor) Start(db DB) {
	d.DB = db
	d.Events = make(chan entities.Event)
	d.Target = &DataTarget{}
	d.Track = &DataTrack{}

	d.Target.Start(db, d.Events)
	d.Track.Start(db, d.Events)
}
