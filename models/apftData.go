package models

//APFTData Model representing apft data
type APFTData struct {
	PercentFail     float32 `json:"percentfail" bson:"percentfail"`
	PercentPassed   float32 `json:"percentpassed" bson:"percentpassed"`
	PercentNotTaken float32 `json:"percentnottaken" bson:"percentnottaken"`
}
