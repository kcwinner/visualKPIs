package routers

import (
	"github.com/gorilla/mux"
	"github.com/kcwinner/visualKPIs/controllers"
)

func setSSDRoutes(router *mux.Router) *mux.Router {
	ssdRouter := mux.NewRouter().StrictSlash(false)
	ssdRouter.HandleFunc("/api/v1/ssd", controllers.GetSSDData).Methods("GET")

	router.PathPrefix("/api/v1/ssd").Handler(ssdRouter)

	return router
}
