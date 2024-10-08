package repositories

import (
	"go-keycloak-sample/src/config"
	"go-keycloak-sample/src/domains"
	"go-keycloak-sample/src/errors"
)

func SaveEvent(event *domains.Event) (*domains.Event, *errors.HttpError) {
	e := config.Database.Create(&event)

	if e.Error != nil {
		return nil, errors.DataAccessLayerError(e.Error.Error())
	}

	return event, nil
}

func FindOneEventById(id int) *domains.Event {
	var event domains.Event

	config.Database.First(&event, id)

	return &event
}

func FindAllEvents() []domains.Event {
	var events []domains.Event
	config.Database.Find(&events)

	return events
}
