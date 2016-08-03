package routers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kcwinner/visualKPIs/controllers"
	"github.com/kcwinner/visualKPIs/handlers"
	"github.com/kcwinner/visualKPIs/models"
	"github.com/kcwinner/visualKPIs/viewModels"
)

func setSoldiersRoutes(router *mux.Router) *mux.Router {
	soldiersRouter := mux.NewRouter().StrictSlash(false)
	soldiersRouter.HandleFunc("/api/v1/soldiers/{id}", controllers.GetSoldierByID).Methods("GET")
	soldiersRouter.HandleFunc("/api/v1/soldiers", controllers.GetSoldiers).Methods("GET")
	soldiersRouter.HandleFunc("/api/v1/soldiers", controllers.AddSoldier).Methods("POST")
	soldiersRouter.HandleFunc("/api/v1/soldiers", controllers.UpdateSoldier).Methods("PUT")
	soldiersRouter.HandleFunc("/api/v1/soldiers/{id}", controllers.DeleteSoldierByID).Methods("DELETE")

	router.PathPrefix("/api/v1/soldiers").Handler(soldiersRouter)

	soldiersWebRouter := mux.NewRouter().StrictSlash(false)
	soldiersWebRouter.HandleFunc("/soldiers", soldiersHandler)
	soldiersWebRouter.HandleFunc("/soldiers/{id}", soldierHandler)
	soldiersWebRouter.HandleFunc("/soldiers/{id}/edit", editSoldierHandler)
	router.PathPrefix("/soldiers").Handler(soldiersWebRouter)

	return router
}

func soldiersHandler(w http.ResponseWriter, r *http.Request) {
	dataTemp := &handlers.DataTemplateHandler{
		Filename: "soldiers.html",
	}

	client := &http.Client{}

	req, err := buildGetRequest("http://localhost:8080/api/v1/soldiers")
	if err != nil {
		log.Println(err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	var soldiers []models.Soldier
	json.Unmarshal(body, &soldiers)

	vm := viewModels.SoldiersViewModel{
		Soldiers: soldiers,
	}

	dataTemp.RenderTemplate(w, "base", vm)
}

func soldierHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	dataTemp := &handlers.DataTemplateHandler{
		Filename: "soldier.html",
	}

	client := &http.Client{}

	req, err := buildGetRequest("http://localhost:8080/api/v1/soldiers/" + id)
	if err != nil {
		log.Println(err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	var soldier models.Soldier
	json.Unmarshal(body, &soldier)

	vm := viewModels.SoldierViewModel{
		Soldier: soldier,
	}

	dataTemp.RenderTemplate(w, "base", vm)
}

func editSoldierHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	dataTemp := &handlers.DataTemplateHandler{
		Filename: "editSoldier.html",
	}

	client := &http.Client{}

	req, err := buildGetRequest("http://localhost:8080/api/v1/soldiers/" + id)
	if err != nil {
		log.Println(err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	var soldier models.Soldier
	json.Unmarshal(body, &soldier)

	vm := viewModels.SoldierViewModel{
		Soldier: soldier,
	}

	dataTemp.RenderTemplate(w, "base", vm)
}
