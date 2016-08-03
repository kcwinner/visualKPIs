package routers

import (
	"github.com/gorilla/mux"
	"github.com/kcwinner/visualKPIs/controllers"
	"github.com/kcwinner/visualKPIs/handlers"
)

func setImportRoutes(router *mux.Router) *mux.Router {
	importRouter := mux.NewRouter().StrictSlash(false)
	importRouter.HandleFunc("/api/v1/import", controllers.ImportData).Methods("POST")

	router.PathPrefix("/api/v1/import").Handler(importRouter)

	importWebRouter := mux.NewRouter().StrictSlash(false)
	importWebRouter.Handle("/import", &handlers.TemplateHandler{
		Filename: "import.html",
		UseBase:  true,
	})

	router.PathPrefix("/import").Handler(importWebRouter)

	return router
}
