package api

// import gorilla mux
import (
	"dreamt/pkg/api/models"
	"dreamt/pkg/controller"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/mux"
)

type API struct {
	router     *mux.Router
	httpClient *http.Client
	controller *controller.Controller
	app        *fiber.App
	webApp     models.WebApp
	addr       string
}

func NewAPI(ctr *controller.Controller, addr string, webApp models.WebApp, httpClient *http.Client) *API {
	api := API{
		addr:       addr,
		webApp:     webApp,
		httpClient: httpClient,
		controller: ctr,
		app:        fiber.New(),
		router:     mux.NewRouter(),
	}

	handleRoutes(&api)

	return &api
}

func handleRoutes(api *API) {
	// add all routes to the router
	for _, route := range getRoutes(api) {
		switch api.webApp {
		case models.Fiber:
			fiberFormattedPath := strings.ReplaceAll(route.Path, "}", "")
			fiberFormattedPath = strings.ReplaceAll(fiberFormattedPath, "{", ":")

			api.app.Add(route.Method, fiberFormattedPath, route.FiberHandlerFunc)
		default:
			api.router.
				Methods(route.Method).
				Path(route.Path).
				Name(route.Name).
				Handler(route.HandlerFunc)
		}
	}
}

func (api *API) Run() error {
	var err error
	switch api.webApp {
	case models.Fiber:
		err = api.app.Listen(api.addr)
	default:
		err = http.ListenAndServe(api.addr, api.router)
	}
	return err
}
