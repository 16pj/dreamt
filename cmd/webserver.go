package main

import (
	"dreamt/pkg/api"
	"dreamt/pkg/api/models"
	"dreamt/pkg/app"
	"dreamt/pkg/controller"
	"dreamt/pkg/persistence/postgres"
	"log"
)

func main() {
	dbController := postgres.NewPGController("")
	ctr := controller.NewController(dbController)
	webApp := models.Fiber
	// webApp = models.GorillaMux
	myAPI := api.NewAPI(&ctr, nil)
	app := app.AppFactory(myAPI, webApp, ":8080")

	// start command line interface
	// start http server
	log.Fatal(app.Run())
}
