package monitor

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DataMonitor ...
type DataMonitor struct {
	DB     DB
	Events chan Event
	target *mgo.Collection
}

// GetTargetByURL ...
func (d *DataMonitor) GetTargetByURL(url string) *Target {
	var target Target
	err := d.target.Find(bson.M{"url": url}).One(&target)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}

		log.Printf("got an error finding a doc %v\n", err)
	}

	return &target
}

// GetTargetByID ...
func (d *DataMonitor) GetTargetByID(id string) *Target {
	_id := bson.ObjectIdHex(id)
	var target Target

	err := d.target.Find(bson.M{"_id": _id}).One(&target)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}

		log.Printf("got an error finding a doc %v\n", err)
	}

	return &target
}

// AddTarget ...
func (d *DataMonitor) AddTarget(url string) {
	target := d.GetTargetByURL(url)
	if target != nil {
		return
	}

	log.Printf("Adding target %s", url)

	doc := Target{ID: bson.NewObjectId(), URL: url, Status: StatusNone}
	if err := d.target.Insert(doc); err != nil {
		log.Printf("Can't insert document: %v\n", err)
	}

	go func() { d.Events <- Event{Event: Added, Target: &doc} }()
}

// RemoveTarget ...
func (d *DataMonitor) RemoveTarget(url string) {
	target := d.GetTargetByURL(url)
	if target == nil {
		log.Printf("Can't find document with url: %s\n", url)
		return
	}

	log.Printf("Removing url %s", url)

	err := d.target.Remove(bson.M{"url": url})
	if err != nil {
		log.Printf("Can't delete document: %v\n", err)
	}

	go func() { d.Events <- Event{Event: Removed, Target: target} }()
}

// RemoveTargetByID ...
func (d *DataMonitor) RemoveTargetByID(id string) {
	target := d.GetTargetByID(id)
	if target == nil {
		log.Printf("Can't find document with id: %s\n", id)
		return
	}

	log.Printf("Removing url %s", target.URL)

	err := d.target.Remove(bson.M{"url": target.URL})
	if err != nil {
		log.Printf("Can't delete document: %v\n", err)
	}

	go func() { d.Events <- Event{Event: Removed, Target: target} }()
}

// UpdateStatusByURL ...
func (d *DataMonitor) UpdateStatusByURL(url string, status string) {
	target := d.GetTargetByURL(url)
	if target != nil {
		log.Printf("Can't find document with url: %s\n", url)
		return
	}

	log.Printf("Updating url %s to status %s", url, status)
	err := d.target.Update(bson.M{"url": url}, bson.M{"status": status})
	if err != nil {
		log.Printf("Can't update document: %v\n", err)
	}

	go func() { d.Events <- Event{Event: Updated, Target: target} }()
}

// UpdateTarget ...
func (d *DataMonitor) UpdateTarget(id string, data Target) {
	target := d.GetTargetByID(id)
	if target == nil {
		log.Printf("Can't find document with id: %s\n", id)
		return
	}

	log.Printf("Updating url %s to status %s", target.URL, target.Status)
	err := d.target.Update(bson.M{"_id": target.ID}, bson.M{"url": data.URL, "status": data.Status})
	if err != nil {
		log.Printf("Can't update document: %v\n", err)
	}

	go func() { d.Events <- Event{Event: Updated, Target: target} }()
}

// GetAllURLS ...
func (d *DataMonitor) GetAllURLS() []string {
	urls := []string{}

	targets := d.GetAllTargets()

	for _, target := range targets {
		urls = append(urls, target.URL)
	}

	return urls
}

// GetAllTargets ...
func (d *DataMonitor) GetAllTargets() []Target {
	targets := []Target{}

	err := d.target.Find(nil).All(&targets)
	if err != nil {
		log.Printf("got an error finding a doc %v\n", err)
	}

	return targets
}

// Start ...
func (d *DataMonitor) Start(db DB) {
	d.DB = db
	d.Events = make(chan Event)
	d.target = d.DB.Session.DB(d.DB.DBName).C("target")
}
