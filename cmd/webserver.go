package main

import (
	"dreamt/pkg/api"
	"dreamt/pkg/controller"
	"dreamt/pkg/persistence/postgres"
	"net/http"
)

func main() {
	dbController := postgres.NewPGController("")
	ctr := controller.NewController(dbController)
	api := api.NewAPI(&ctr, nil)
	// start http server
	http.ListenAndServe(":8080", api.Router)
}
