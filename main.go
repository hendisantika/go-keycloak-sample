package main

import (
	"github.com/gorilla/mux"
	"go-keycloak-sample/src/config"
	"go-keycloak-sample/src/controllers"
	"go-keycloak-sample/src/services"
	"log"
	"net/http"
)

func run() {
	config.DbConnect()
	services.InitializeOauthServer()
	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)

	registerRoutes(router)

	log.Fatal(http.ListenAndServe(":8081", router))
}

func registerRoutes(router *mux.Router) {
	registerControllerRoutes(controllers.EventController{}, router)
}

func registerControllerRoutes(controller controllers.Controller, router *mux.Router) {
	controller.RegisterRoutes(router)
}
