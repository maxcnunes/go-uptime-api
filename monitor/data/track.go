package data

import (
	"log"

	"github.com/maxcnunes/monitor-api/monitor/entities"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DataTrack ...
type DataTrack struct {
	collection *mgo.Collection
	events     chan entities.Event
}

// GetAll ...
func (d *DataTrack) GetAll() []entities.Track {
	tracks := []entities.Track{}

	err := d.collection.Find(nil).All(&tracks)
	if err != nil {
		log.Printf("got an error finding a doc %v\n", err)
	}

	return tracks
}

// Create ...
func (d *DataTrack) Create(target entities.Target, status string) *entities.Track {
	doc := entities.Track{ID: bson.NewObjectId(), TargetID: target.ID, Status: status}
	if err := d.collection.Insert(doc); err != nil {
		log.Printf("Can't insert document: %v\n", err)
	}

	return &doc
}

// Start ...
func (d *DataTrack) Start(db DB, events chan entities.Event) {
	d.collection = db.Session.DB(db.DBName).C("track")
	d.events = events
}
