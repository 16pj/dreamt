package api

// import gorilla mux
import (
	"dreamt/pkg/controller"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	Router     *mux.Router
	HTTPClient *http.Client
	Controller *controller.Controller
}

func NewAPI(ctr *controller.Controller, httpClient *http.Client) *API {
	router := mux.NewRouter()
	api := API{
		Router:     router,
		HTTPClient: httpClient,
		Controller: ctr,
	}

	// add all routes to the router
	for _, route := range getRoutes(&api) {
		api.Router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return &api
}
