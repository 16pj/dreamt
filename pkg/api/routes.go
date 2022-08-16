package api

import "net/http"

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

func getRoutes(api *API) []Route {
	return []Route{
		{
			Name:        "GetDreams",
			Method:      "GET",
			Path:        "/dreams",
			HandlerFunc: api.GetDreams,
		},
		{
			Name:        "GetDream",
			Method:      "GET",
			Path:        "/dreams/{id}",
			HandlerFunc: api.GetDream,
		},
		{
			Name:        "GetInterpretation",
			Method:      "GET",
			Path:        "/interpret/{keyword}",
			HandlerFunc: api.GetInterpretation,
		},
		{
			Name:        "GetKeywords",
			Method:      "GET",
			Path:        "/keywords",
			HandlerFunc: api.GetKeywords,
		},
		{
			Name:        "CreateDream",
			Method:      "POST",
			Path:        "/dream",
			HandlerFunc: api.CreateDream,
		},
		{
			Name:        "DeleteDream",
			Method:      "DELETE",
			Path:        "/dreams/{id}",
			HandlerFunc: api.DeleteDream,
		},
	}
}
