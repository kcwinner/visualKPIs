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

func setSlideshowRoutes(router *mux.Router) *mux.Router {
	slideshowRouter := mux.NewRouter().StrictSlash(false)
	slideshowRouter.HandleFunc("/slideshow", slideshowHandler)

	router.PathPrefix("/slideshow").Handler(slideshowRouter)

	return router
}

func slideshowHandler(w http.ResponseWriter, r *http.Request) {
	dataTemp := &handlers.DataTemplateHandler{
		Filename: "slideshow.html",
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

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	var soldiers []models.Soldier
	json.Unmarshal(body, &soldiers)

	response.Body.Close()

	req, err = buildGetRequest("http://localhost:8080/api/v1/apft")
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

	vm := viewModels.SlideshowViewModel{
		Soldiers: soldiers,
		APFTData: apftData,
		SSDData:  ssdData,
	}

	dataTemp.RenderTemplate(w, "base", vm)
}
