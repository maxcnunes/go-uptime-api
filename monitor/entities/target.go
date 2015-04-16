package entities

import "gopkg.in/mgo.v2/bson"

// Target data structure
type Target struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	URL    string        `bson:"url" json:"url"`
	Status int           `bson:"status" json:"status"`
	Emails []string      `bson:"emails" json:"emails"`
}
