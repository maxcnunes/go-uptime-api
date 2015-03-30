package entities

import "gopkg.in/mgo.v2/bson"

// Track ...
type Track struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	TargetID bson.ObjectId `bson:"targetId" json:"targetId"`
	Status   string        `bson:"status" json:"status"`
}
