package routers

import (
	"github.com/gorilla/mux"
	"github.com/kcwinner/visualKPIs/handlers"
)

func setWebRoutes(router *mux.Router) *mux.Router {
	loginRouter := mux.NewRouter().StrictSlash(false)
	loginRouter.Handle("/login", &handlers.TemplateHandler{
		Filename: "login.html",
		UseBase:  false,
	})

	router.PathPrefix("/login").Handler(loginRouter)

	return router
}
