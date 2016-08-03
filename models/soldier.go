package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Soldier Model representing a soldier
type Soldier struct {
	ID                 bson.ObjectId `json:"id" bson:"_id,omitempty"`
	FirstName          string        `json:"firstname" bson:"firstname"`
	LastName           string        `json:"lastname" bson:"lastname"`
	Rank               string        `json:"rank" bson:"rank"`
	MOS                string        `json:"mos" bson:"mos"`
	Section            string        `json:"section" bson:"section"`
	CivilianEmployment string        `json:"civilianemployment" bson:"civilianemployment"`
	ETSDate            time.Time     `json:"etsdate" bson:"etsdate"`
	NCOERDate          time.Time     `json:"ncoerdate" bson:"ncoerdate"`
	APFTPass           bool          `json:"apftpass" bson:"apftpass"`
	APFTScore          int           `json:"apftscore" bson:"apftscore"`
	SSD                int           `json:"ssd" bson:"ssd"`
	HTWT               string        `json:"htwt" bson:"htwt"`
	Image              string        `json:"image" bson:"image"`
}
