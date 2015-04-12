package entities

import "gopkg.in/mgo.v2/bson"

// Target ...
type Target struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	URL    string        `bson:"url" json:"url"`
	Status int           `bson:"status" json:"status"`
}
