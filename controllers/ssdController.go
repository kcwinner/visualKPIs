package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/kcwinner/visualKPIs/common"
	"github.com/kcwinner/visualKPIs/data"
	"github.com/kcwinner/visualKPIs/models"
)

//GetSSDData HTTP Get - /api/v1/ssd
func GetSSDData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	session := data.GetDb()
	defer session.Close()

	coll := session.DB(common.AppConfig.Database).C("Soldiers")
	iter := coll.Find(nil).Iter()

	result := models.Soldier{}

	totalPercent := 0
	soldierCount := 0

	for iter.Next(&result) {
		totalPercent += result.SSD
		soldierCount++
	}

	var ssdData models.SSDData

	finalPercent := float32(totalPercent) / float32(soldierCount*100) * 100

	ssdData.PercentFinished = finalPercent
	ssdData.PercentUnfinished = 100 - finalPercent

	ssdJSON, err := json.Marshal(ssdData)

	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Error getting SSD data.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(ssdJSON)
}
