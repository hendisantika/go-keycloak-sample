package controllers

import (
	"github.com/gorilla/mux"
	"go-keycloak-sample/src/services"
	"net/http"
)

type EventController struct {
}

func (t EventController) RegisterRoutes(router *mux.Router) {
	router.Handle("/events", services.Protect(http.HandlerFunc(services.CreateEvent))).Methods("POST")
	router.Handle("/events/{id}", services.Protect(http.HandlerFunc(services.GetOneEvent))).Methods("GET")
	router.Handle("/events", services.Protect(http.HandlerFunc(services.AllEvents))).Methods("GET")
}
