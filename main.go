package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-keycloak-sample/src/config"
	"go-keycloak-sample/src/controllers"
	"go-keycloak-sample/src/services"
	"log"
	"net/http"
)

func main() {
	run()
}

func run() {
	config.DbConnect()
	services.InitializeOauthServer()
	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)

	registerRoutes(router)

	fmt.Println("Hello .... Start the service")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func registerRoutes(router *mux.Router) {
	registerControllerRoutes(controllers.EventController{}, router)
}

func registerControllerRoutes(controller controllers.Controller, router *mux.Router) {
	controller.RegisterRoutes(router)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
