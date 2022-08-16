package app

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	Name             string
	Method           string
	Path             string
	HandlerFunc      http.HandlerFunc
	FiberHandlerFunc fiber.Handler
}

func getRoutes(app *App) []Route {
	return []Route{
		{
			Name:             "GetDreams",
			Method:           "GET",
			Path:             "/dreams",
			HandlerFunc:      app.gorillaApp.GetDreams,
			FiberHandlerFunc: app.fiberApp.FGetDreams,
		},
		{
			Name:             "GetDream",
			Method:           "GET",
			Path:             "/dreams/{id}",
			HandlerFunc:      app.gorillaApp.GetDream,
			FiberHandlerFunc: app.fiberApp.FGetDream,
		},
		{
			Name:             "GetInterpretation",
			Method:           "GET",
			Path:             "/interpret/{keyword}",
			HandlerFunc:      app.gorillaApp.GetInterpretation,
			FiberHandlerFunc: app.fiberApp.FGetInterpretation,
		},
		{
			Name:             "GetKeywords",
			Method:           "GET",
			Path:             "/keywords",
			HandlerFunc:      app.gorillaApp.GetKeywords,
			FiberHandlerFunc: app.fiberApp.FGetKeywords,
		},
		{
			Name:             "CreateDream",
			Method:           "POST",
			Path:             "/dream",
			HandlerFunc:      app.gorillaApp.CreateDream,
			FiberHandlerFunc: app.fiberApp.FCreateDream,
		},
		{
			Name:             "DeleteDream",
			Method:           "DELETE",
			Path:             "/dreams/{id}",
			HandlerFunc:      app.gorillaApp.DeleteDream,
			FiberHandlerFunc: app.fiberApp.FDeleteDream,
		},
	}
}
