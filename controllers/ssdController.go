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
	var ssdData models.SSDData

	session := data.GetDb()
	defer session.Close()

	coll := session.DB(common.AppConfig.Database).C("SSD")
	err := coll.Find("").One(&ssdData)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Error getting SSD data.")
		return
	}

	ssdJSON, err := json.Marshal(ssdData)

	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Error getting SSD data.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(ssdJSON)
}
