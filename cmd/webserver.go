package main

import (
	"dreamt/pkg/api"
	"dreamt/pkg/api/models"
	"dreamt/pkg/controller"
	"dreamt/pkg/persistence/postgres"
	"log"
)

func main() {
	dbController := postgres.NewPGController("")
	ctr := controller.NewController(dbController)
	webApp := models.Fiber
	// webApp = models.GorillaMux
	api := api.NewAPI(&ctr, ":8080", webApp, nil)

	// start command line interface
	// start http server
	log.Fatal(api.Run())
}
