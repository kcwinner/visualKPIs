package routers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kcwinner/visualKPIs/handlers"
	"github.com/kcwinner/visualKPIs/models"
	"github.com/kcwinner/visualKPIs/viewModels"
)

func setHomeRoutes(router *mux.Router) *mux.Router {
	fs := http.FileServer(http.Dir("assets"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	homeRouter := mux.NewRouter().StrictSlash(false)
	homeRouter.HandleFunc("/home", homeHandler)

	// homeRouter.Handle("/home", &handlers.TemplateHandler{
	// 	Filename: "home.html",
	// 	UseBase:  true,
	// })

	router.PathPrefix("/home").Handler(homeRouter)

	return router
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	dataTemp := &handlers.DataTemplateHandler{
		Filename: "home.html",
	}

	client := &http.Client{}

	req, err := buildGetRequest("http://localhost:8080/api/v1/apft")
	if err != nil {
		log.Println(err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	var apftData models.APFTData
	json.Unmarshal(body, &apftData)

	response.Body.Close()

	req, err = buildGetRequest("http://localhost:8080/api/v1/ssd")
	if err != nil {
		log.Println(err)
	}

	response, err = client.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	var ssdData models.SSDData
	json.Unmarshal(body, &ssdData)

	response.Body.Close()

	req, err = buildGetRequest("http://localhost:8080/api/v1/soldiers")
	if err != nil {
		log.Println(err)
	}

	response, err = client.Do(req)
	if err != nil {
		log.Println(err)
	}

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	var soldiers []models.Soldier
	json.Unmarshal(body, &soldiers)

	response.Body.Close()

	vm := viewModels.HomeViewModel{
		APFTData: apftData,
		SSDData:  ssdData,
		Soldiers: soldiers,
	}

	dataTemp.RenderTemplate(w, "base", vm)
}
