package entities

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Track data structure
type Track struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	TargetID  bson.ObjectId `bson:"targetId" json:"targetId"`
	Status    int           `bson:"status" json:"status"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
}
