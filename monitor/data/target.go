package data

import (
	"log"

	"github.com/maxcnunes/monitor-api/monitor/entities"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DataTarget ...
type DataTarget struct {
	collection *mgo.Collection
	events     chan entities.Event
}

// FindOneByURL ...
func (d *DataTarget) FindOneByURL(url string) *entities.Target {
	var target entities.Target
	err := d.collection.Find(bson.M{"url": url}).One(&target)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}

		log.Printf("got an error finding a doc %v\n", err)
	}

	return &target
}

// FindOneByID ...
func (d *DataTarget) FindOneByID(id string) *entities.Target {
	_id := bson.ObjectIdHex(id)
	var target entities.Target

	err := d.collection.Find(bson.M{"_id": _id}).One(&target)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}

		log.Printf("got an error finding a doc %v\n", err)
	}

	return &target
}

// Create ...
func (d *DataTarget) Create(url string) *entities.Target {
	target := d.FindOneByURL(url)
	if target != nil {
		return target
	}

	log.Printf("Adding target %s", url)

	doc := entities.Target{ID: bson.NewObjectId(), URL: url, Status: 0}
	if err := d.collection.Insert(doc); err != nil {
		log.Printf("Can't insert document: %v\n", err)
	}

	go func() { d.events <- entities.Event{Event: entities.Added, Target: &doc} }()

	return &doc
}

// Remove ...
func (d *DataTarget) Remove(url string) {
	target := d.FindOneByURL(url)
	if target == nil {
		log.Printf("Can't find document with url: %s\n", url)
		return
	}

	log.Printf("Removing url %s", url)

	err := d.collection.Remove(bson.M{"url": url})
	if err != nil {
		log.Printf("Can't delete document: %v\n", err)
	}

	go func() { d.events <- entities.Event{Event: entities.Removed, Target: target} }()
}

// RemoveByID ...
func (d *DataTarget) RemoveByID(id string) {
	target := d.FindOneByID(id)
	if target == nil {
		log.Printf("Can't find document with id: %s\n", id)
		return
	}

	log.Printf("Removing url %s", target.URL)

	err := d.collection.Remove(bson.M{"url": target.URL})
	if err != nil {
		log.Printf("Can't delete document: %v\n", err)
	}

	go func() { d.events <- entities.Event{Event: entities.Removed, Target: target} }()
}

// UpdateStatusByURL ...
func (d *DataTarget) UpdateStatusByURL(url string, status string) {
	target := d.FindOneByURL(url)
	if target != nil {
		log.Printf("Can't find document with url: %s\n", url)
		return
	}

	log.Printf("Updating url %s to status %s", url, status)
	err := d.collection.Update(bson.M{"url": url}, bson.M{"status": status})
	if err != nil {
		log.Printf("Can't update document: %v\n", err)
	}

	go func() { d.events <- entities.Event{Event: entities.Updated, Target: target} }()
}

// Update ...
func (d *DataTarget) Update(id string, data entities.Target) {
	target := d.FindOneByID(id)
	if target == nil {
		log.Printf("Can't find document with id: %s\n", id)
		return
	}

	log.Printf("Updating url %s to status %d", target.URL, target.Status)
	err := d.collection.Update(bson.M{"_id": target.ID}, bson.M{"url": data.URL, "status": data.Status})
	if err != nil {
		log.Printf("Can't update document: %v\n", err)
	}

	go func() { d.events <- entities.Event{Event: entities.Updated, Target: target} }()
}

// GetAllURLS ...
func (d *DataTarget) GetAllURLS() []string {
	urls := []string{}

	targets := d.GetAll()

	for _, target := range targets {
		urls = append(urls, target.URL)
	}

	return urls
}

// GetAll ...
func (d *DataTarget) GetAll() []entities.Target {
	targets := []entities.Target{}

	err := d.collection.Find(nil).All(&targets)
	if err != nil {
		log.Printf("got an error finding a doc %v\n", err)
	}

	return targets
}

// Start ...
func (d *DataTarget) Start(db DB, events chan entities.Event) {
	d.collection = db.Session.DB(db.DBName).C("target")
	d.events = events
}
