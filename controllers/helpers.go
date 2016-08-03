package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

func sendError(w http.ResponseWriter, err error, status int, message string) {
	log.Println(err)
	resp, _ := json.Marshal(message)

	w.WriteHeader(status)
	w.Write(resp)
}

func jsonResponse(w http.ResponseWriter, response interface{}) {
	json, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
