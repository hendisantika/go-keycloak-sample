package services

import (
	"encoding/json"
	"fmt"
	"go-keycloak-sample/src/domains"
	"go-keycloak-sample/src/errors"
	"go-keycloak-sample/src/repositories"
	"io/ioutil"
	"net/http"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent domains.Event
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)

	ev, httpErr := repositories.SaveEvent(&newEvent)

	if httpErr != nil {
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(errors.UnauthorizedError())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&ev)
}
