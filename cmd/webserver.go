package main

import (
	"dreamt/pkg/api"
	"dreamt/pkg/api/models"
	"dreamt/pkg/app"
	"dreamt/pkg/controller"
	"dreamt/pkg/persistence"
	"dreamt/pkg/persistence/mongo"
	"dreamt/pkg/persistence/postgres"
	"fmt"
	"log"
	"time"
)

func main() {

	defer func() {
		// recover
		if r := recover(); r != nil {
			fmt.Println("recovering: ", r)
		}
	}()

	dbType := persistence.PG
	dbType = persistence.MG
	dbController := getDBController(dbType)
	ctr := controller.NewController(dbController)
	webApp := models.Fiber
	// webApp = models.GorillaMux
	myAPI := api.NewAPI(&ctr, nil)
	app := app.AppFactory(myAPI, webApp, ":8080")

	// start command line interface
	// start http server
	log.Fatal(app.Run())
}

func getDBController(dbtype persistence.Database) (dbController persistence.DatabaseController) {
	switch dbtype {
	case persistence.MG:
		dbController = mongo.NewMGController("", time.Second*30)
	default:
		dbController = postgres.NewPGController("")
	}
	return
}
