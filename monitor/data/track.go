package data

import (
	"log"
	"time"

	"github.com/maxcnunes/go-uptime-api/monitor/entities"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DataTrack is the data configuration related to Track collection
type DataTrack struct {
	collection *mgo.Collection
	events     chan entities.Event
}

// Find all tracks
func (d *DataTrack) Find(targetID string) []entities.Track {
	tracks := []entities.Track{}

	query := bson.M{}
	if targetID != "" {
		query["targetId"] = bson.ObjectIdHex(targetID)
	}

	err := d.collection.Find(query).Limit(50).Sort("-createdAt").All(&tracks)
	if err != nil {
		log.Printf("got an error finding a doc %v\n", err)
	}

	return tracks
}

// Create a new track
func (d *DataTrack) Create(target entities.Target, status int) *entities.Track {
	doc := entities.Track{
		ID:        bson.NewObjectId(),
		TargetID:  target.ID,
		Status:    status,
		CreatedAt: time.Now(),
	}
	if err := d.collection.Insert(doc); err != nil {
		log.Printf("Can't insert document: %v\n", err)
	}

	return &doc
}

// RemoveByTargetID removes a track by the targetId field
func (d *DataTrack) RemoveByTargetID(id string) {
	err := d.collection.Remove(bson.M{"targetId": bson.ObjectIdHex(id)})
	if err != nil {
		log.Printf("Can't delete document: %v\n", err)
	}
}

// Start a new instance of data track
func (d *DataTrack) Start(db DB, events chan entities.Event) {
	d.collection = db.Session.DB(db.DBName).C("track")
	d.events = events
}
