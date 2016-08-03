package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/kcwinner/visualKPIs/data"
	"github.com/kcwinner/visualKPIs/models"
)

//GetPTData HTTP Get - /api/v1/apft
func GetPTData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var apftData models.APFTData

	session := data.GetDb()
	defer session.Close()

	coll := session.DB("goDb").C("APFT")
	err := coll.Find("").One(&apftData)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Error getting APFT data.")
		return
	}

	apftJSON, err := json.Marshal(apftData)

	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Error getting APFT data.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(apftJSON)
}
