package routers

import (
	"github.com/gorilla/mux"
	"github.com/kcwinner/visualKPIs/controllers"
)

func setAPFTRoutes(router *mux.Router) *mux.Router {
	apftRouter := mux.NewRouter().StrictSlash(false)
	apftRouter.HandleFunc("/api/v1/apft", controllers.GetPTData).Methods("GET")

	router.PathPrefix("/api/v1/apft").Handler(apftRouter)

	return router
}
