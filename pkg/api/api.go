package api

// import gorilla mux
import (
	"dreamt/pkg/controller"
	"net/http"
)

type API struct {
	httpClient *http.Client
	controller *controller.Controller
	gorillaApp GorillaAPI
	addr       string
}

func NewAPI(ctr *controller.Controller, httpClient *http.Client) *API {
	return &API{
		httpClient: httpClient,
		controller: ctr,
	}
}
