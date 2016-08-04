package routers

import (
	"github.com/gorilla/mux"
)

//InitRoutes Initializes all the routes for the api
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	router = setAPFTRoutes(router)
	router = setHomeRoutes(router)
	router = setImportRoutes(router)
	router = setSlideshowRoutes(router)
	router = setSoldiersRoutes(router)
	router = setSSDRoutes(router)
	router = setWebRoutes(router)

	return router
}
