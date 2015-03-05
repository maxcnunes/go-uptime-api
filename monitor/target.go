package monitor

import (
	"gopkg.in/mgo.v2/bson"
)

// Target ...
type Target struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	URL    string        `bson:"url" json:"url"`
	Status string        `bson:"status" json:"status"`
}

// Status
const (
	StatusUp   = "up"
	StatusDown = "down"
	StatusNone = "none"
)
