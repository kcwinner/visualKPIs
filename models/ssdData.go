package models

//SSDData Model representing ssd data
type SSDData struct {
	PercentFinished   float32 `json:"percentfinished" bson:"percentfinished"`
	PercentUnfinished float32 `json:"percentunfinished" bson:"percentunfinished"`
}
