package models

import (
	"gopkg.in/mgo.v2/bson"
)

//APFTData Model representing apft data
type APFTData struct {
	ID              bson.ObjectId `json:"id" bson:"_id,omitempty"`
	PercentFail     float32       `json:"percentfail" bson:"percentfail"`
	PercentPassed   float32       `json:"percentpassed" bson:"percentpassed"`
	PercentNotTaken float32       `json:"percentnottaken" bson:"percentnottaken"`
}
