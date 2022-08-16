package controller

import "dreamt/pkg/persistence/postgres"

type Controller struct {
	DBController postgres.PGController
}

func NewController(dbController postgres.PGController) Controller {
	return Controller{
		DBController: dbController,
	}
}
