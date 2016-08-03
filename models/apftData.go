package models

import (
	"gopkg.in/mgo.v2/bson"
)

//APFTData Model representing apft data
type APFTData struct {
	ID              bson.ObjectId `json:"id" bson:"_id,omitempty"`
	PercentFail     int           `json:"percentfail" bson:"percentfail"`
	PercentPassed   int           `json:"percentpassed" bson:"percentpassed"`
	PercentNotTaken int           `json:"percentnottaken" bson:"percentnottaken"`
}
