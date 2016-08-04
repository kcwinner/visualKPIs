package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kcwinner/visualKPIs/common"
	"github.com/kcwinner/visualKPIs/data"
	"github.com/kcwinner/visualKPIs/models"
	"gopkg.in/mgo.v2/bson"
)

//GetSoldierByID HTTP Get - /api/v1/soldiers/{id}
func GetSoldierByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])

	session := data.GetDb()
	defer session.Close()

	var soldier models.Soldier

	coll := session.DB(common.AppConfig.Database).C("Soldiers")
	err := coll.FindId(id).One(&soldier)

	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Error getting soldier.")
		return
	}

	soldierJSON, err := json.Marshal(soldier)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Error getting soldier.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(soldierJSON)
}

//GetSoldiers HTTP Get - /api/v1/soldiers
func GetSoldiers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var soldiers []models.Soldier

	session := data.GetDb()
	defer session.Close()

	coll := session.DB(common.AppConfig.Database).C("Soldiers")
	iter := coll.Find(nil).Sort("lastname").Iter()
	result := models.Soldier{}
	for iter.Next(&result) {
		soldiers = append(soldiers, result)
	}

	soldiersJSON, err := json.Marshal(soldiers)

	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Error getting soldiers.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(soldiersJSON)
}

//AddSoldier HTTP Post - /api/v1/soldiers
func AddSoldier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var soldier models.Soldier

	err := json.NewDecoder(r.Body).Decode(&soldier)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Incorrect data format.")
		return
	}

	session := data.GetDb()
	defer session.Close()

	soldier.ID = bson.NewObjectId()

	coll := session.DB(common.AppConfig.Database).C("Soldiers")
	err = coll.Insert(soldier)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Error creating soldier.")
		return
	}

	respJSON, _ := json.Marshal("Successfully created soldier.")

	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}

//UpdateSoldier HTTP Put - /api/v1/soldiers
func UpdateSoldier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var soldier models.Soldier

	err := json.NewDecoder(r.Body).Decode(&soldier)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Incorrect data format.")
		return
	}

	session := data.GetDb()
	defer session.Close()

	coll := session.DB(common.AppConfig.Database).C("Soldiers")
	err = coll.UpdateId(soldier.ID, soldier)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Error updating soldier.")
		return
	}

	w.WriteHeader(http.StatusOK)

	respJSON, _ := json.Marshal("Successfully updated soldier.")
	w.Write(respJSON)
}

//DeleteSoldierByID HTTP Delete - /api/v1/soldiers/{id}
func DeleteSoldierByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])

	session := data.GetDb()
	defer session.Close()

	coll := session.DB(common.AppConfig.Database).C("Soldiers")
	err := coll.RemoveId(id)
	if err != nil {
		sendError(w, err, http.StatusInternalServerError, "Error deleting soldier.")
		return
	}

	respJSON, _ := json.Marshal("Successfully deleted soldier.")

	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}
