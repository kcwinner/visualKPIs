package routers

import (
	"github.com/gorilla/mux"
	"github.com/kcwinner/visualKPIs/controllers"
)

func setUsersRoutes(router *mux.Router) *mux.Router {
	userRouter := mux.NewRouter().StrictSlash(false)
	userRouter.HandleFunc("/api/v1/users/login", controllers.HandleLogin).Methods("POST")

	router.PathPrefix("/api/v1/users").Handler(userRouter)

	return router
}
