package controllers

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/kcwinner/visualKPIs/data"
	"github.com/kcwinner/visualKPIs/models"
)

//GetPTData HTTP Get - /api/v1/apft
func GetPTData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var apftData models.APFTData

	session := data.GetDb()
	defer session.Close()

	coll := session.DB("goDb").C("Soldiers")

	passedQuery := bson.M{"apftpass": true}
	untakenQuery := bson.M{"apftscore": 0}

	result := models.Soldier{}

	count, err := coll.Find(nil).Count()
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Error getting apft data.")
	}

	var passedSoldiers []models.Soldier
	iter := coll.Find(passedQuery).Iter()
	for iter.Next(&result) {
		passedSoldiers = append(passedSoldiers, result)
	}

	var untakenSoldiers []models.Soldier
	iter = coll.Find(untakenQuery).Iter()
	for iter.Next(&result) {
		untakenSoldiers = append(untakenSoldiers, result)
	}

	apftData.PercentPassed = float32(len(passedSoldiers)) / float32(count) * 100
	apftData.PercentNotTaken = float32(len(untakenSoldiers)) / float32(count) * 100
	apftData.PercentFail = float32((count - len(untakenSoldiers) - len(passedSoldiers))) / float32(count) * 100

	apftJSON, err := json.Marshal(apftData)

	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Error getting APFT data.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(apftJSON)
}
