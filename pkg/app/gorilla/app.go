package gorilla

import (
	"dreamt/pkg/api"
	rmodels "dreamt/pkg/api/models"
	"dreamt/pkg/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type GorillaAPI struct {
	*mux.Router
	*api.API
}

func NewGorillaAPI(api *api.API) *GorillaAPI {
	return &GorillaAPI{
		mux.NewRouter(),
		api,
	}
}

func sendResp(w http.ResponseWriter, resp rmodels.APIResponse) {
	if resp.Err != nil {
		w.WriteHeader(resp.Status)
		w.Write([]byte(resp.Err.Error()))
		return
	}

	// write response
	if err := json.NewEncoder(w).Encode(resp.Body); err != nil {
		sendResp(w, rmodels.APIResponse{
			Status: http.StatusInternalServerError,
			Err:    err})
	}
}

func (g GorillaAPI) GetDreams(w http.ResponseWriter, r *http.Request) {
	// get dreams from controller
	sendResp(w, g.API.GetDreams())
}

func (g GorillaAPI) GetDream(w http.ResponseWriter, r *http.Request) {
	// get id from url
	vars := mux.Vars(r)
	id := vars["id"]

	// get dreams from controller
	sendResp(w, g.API.GetDream(rmodels.GetDreamRequest{ID: id}))
}

func (g GorillaAPI) GetInterpretation(w http.ResponseWriter, r *http.Request) {
	// get keyword from url
	vars := mux.Vars(r)
	keyword := vars["keyword"]

	// get interpretation from controller
	sendResp(w, g.GetInterpret(rmodels.GetInterpretationRequest{Keyword: keyword}))
}

func (g GorillaAPI) GetKeywords(w http.ResponseWriter, r *http.Request) {
	// get top from query
	limit := r.URL.Query().Get("limit")

	// get keywords from controller
	sendResp(w, g.API.GetKeywords(rmodels.GetKeywordsRequest{Limit: limit}))
}

func (g GorillaAPI) CreateDream(w http.ResponseWriter, r *http.Request) {
	// get dream from body
	var dream models.Dream
	if err := json.NewDecoder(r.Body).Decode(&dream); err != nil {
		sendResp(w, rmodels.APIResponse{Status: http.StatusInternalServerError, Err: err})
		return
	}

	// create dream in controller
	sendResp(w, g.API.CreateDream(rmodels.CreateDreamRequest{Dream: dream}))
}

func (g GorillaAPI) DeleteDream(w http.ResponseWriter, r *http.Request) {
	// get id from url
	vars := mux.Vars(r)
	id := vars["id"]

	// delete dream in controller
	sendResp(w, g.API.DeleteDream(rmodels.DeleteDreamRequest{ID: id}))
}
