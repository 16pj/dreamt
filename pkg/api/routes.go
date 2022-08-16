package api

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

func getRoutes(api *API) []Route {
	return []Route{
		{
			Name:             "GetDreams",
			Method:           "GET",
			Path:             "/dreams",
			HandlerFunc:      api.GetDreams,
			FiberHandlerFunc: api.FGetDreams,
		},
		{
			Name:             "GetDream",
			Method:           "GET",
			Path:             "/dreams/:id",
			HandlerFunc:      api.GetDream,
			FiberHandlerFunc: api.FGetDream,
		},
		{
			Name:             "GetInterpretation",
			Method:           "GET",
			Path:             "/interpret/{keyword}",
			HandlerFunc:      api.GetInterpretation,
			FiberHandlerFunc: api.FGetInterpretation,
		},
		{
			Name:             "GetKeywords",
			Method:           "GET",
			Path:             "/keywords",
			HandlerFunc:      api.GetKeywords,
			FiberHandlerFunc: api.FGetKeywords,
		},
		{
			Name:             "CreateDream",
			Method:           "POST",
			Path:             "/dream",
			HandlerFunc:      api.CreateDream,
			FiberHandlerFunc: api.FCreateDream,
		},
		{
			Name:             "DeleteDream",
			Method:           "DELETE",
			Path:             "/dreams/{id}",
			HandlerFunc:      api.DeleteDream,
			FiberHandlerFunc: api.FDeleteDream,
		},
	}
}
