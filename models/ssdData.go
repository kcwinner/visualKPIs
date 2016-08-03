package models

import (
	"gopkg.in/mgo.v2/bson"
)

//SSDData Model representing ssd data
type SSDData struct {
	ID                bson.ObjectId `json:"id" bson:"_id,omitempty"`
	PercentFinished   int           `json:"percentfinished" bson:"percentfinished"`
	PercentUnfinished int           `json:"percentunfinished" bson:"percentunfinished"`
}
