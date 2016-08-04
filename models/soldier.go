package models

import "gopkg.in/mgo.v2/bson"

//Soldier Model representing a soldier
type Soldier struct {
	ID                 bson.ObjectId `json:"id" bson:"_id,omitempty"`
	FirstName          string        `json:"firstname" bson:"firstname"`
	LastName           string        `json:"lastname" bson:"lastname"`
	Rank               string        `json:"rank" bson:"rank"`
	MOS                string        `json:"mos" bson:"mos"`
	Section            string        `json:"section" bson:"section"`
	CivilianEmployment string        `json:"civilianemployment" bson:"civilianemployment"`
	ETSDate            string        `json:"etsdate" bson:"etsdate"`
	NCOERDate          string        `json:"ncoerdate" bson:"ncoerdate"`
	APFTPass           bool          `json:"apftpass" bson:"apftpass"`
	APFTScore          int           `json:"apftscore" bson:"apftscore"`
	SSD1               int           `json:"ssd1" bson:"ssd1"`
	SSD2               int           `json:"ssd2" bson:"ssd2"`
	SSD3               int           `json:"ssd3" bson:"ssd3"`
	SSD4               int           `json:"ssd4" bson:"ssd4"`
	HTWT               string        `json:"htwt" bson:"htwt"`
	Image              string        `json:"image" bson:"image"`
}
