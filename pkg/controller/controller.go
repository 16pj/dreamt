package controller

import (
	"dreamt/pkg/persistence"
)

type Controller struct {
	DBController persistence.DatabaseController
}

func NewController(dbController persistence.DatabaseController) Controller {
	return Controller{
		DBController: dbController,
	}
}
