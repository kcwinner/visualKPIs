package main

import (
	"log"
	"net/http"

	"github.com/kcwinner/visualKPIs/common"
	"github.com/kcwinner/visualKPIs/data"
	"github.com/kcwinner/visualKPIs/routers"
)

func main() {
	log.Println("Starting Personnel Bio Server...")

	common.StartUp()

	router := routers.InitRoutes()

	dbError := data.DialDb()
	if dbError != nil {
		panic(dbError)
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		data.CloseDb()
	}
}
