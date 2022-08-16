package app

import (
	"dreamt/pkg/api"
	"dreamt/pkg/api/models"
	"net/http"
	"strings"
)

type App struct {
	webApp     models.WebApp
	fiberApp   *api.FiberAPI
	gorillaApp *api.GorillaAPI
	addr       string
}

func AppFactory(myAPI *api.API, webApp models.WebApp, addr string) *App {
	app := App{
		webApp:     webApp,
		addr:       addr,
		fiberApp:   &api.FiberAPI{},
		gorillaApp: &api.GorillaAPI{},
	}

	switch webApp {
	case models.Fiber:
		app.fiberApp = api.NewFiberAPI(myAPI)
	default:
		app.gorillaApp = api.NewGorillaAPI(myAPI)
	}

	return handleRoutes(&app)
}

func handleRoutes(app *App) *App {
	// add all routes to the router
	for _, route := range getRoutes(app) {
		switch app.webApp {
		case models.Fiber:
			fiberFormattedPath := strings.ReplaceAll(route.Path, "}", "")
			fiberFormattedPath = strings.ReplaceAll(fiberFormattedPath, "{", ":")

			app.fiberApp.Add(route.Method, fiberFormattedPath, route.FiberHandlerFunc)
		default:
			app.gorillaApp.
				Methods(route.Method).
				Path(route.Path).
				Name(route.Name).
				Handler(route.HandlerFunc)
		}
	}

	return app
}

func (app *App) Run() error {
	var err error
	switch app.webApp {
	case models.Fiber:
		err = app.fiberApp.Listen(app.addr)
	default:
		err = http.ListenAndServe(app.addr, app.gorillaApp)
	}
	return err
}
