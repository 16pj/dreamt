package api

import (
	"dreamt/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func sendError(w http.ResponseWriter, status int, errMsg string) {
	w.WriteHeader(status)
	w.Write([]byte(errMsg))
}

func sendResp(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(body); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
	}
}

func (a API) GetDreams(w http.ResponseWriter, r *http.Request) {
	// get dreams from controller
	dreams, err := a.Controller.GetDreams()
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// write response
	sendResp(w, dreams)
}

func (a API) GetDream(w http.ResponseWriter, r *http.Request) {
	// get id from url
	vars := mux.Vars(r)
	id := vars["id"]

	// get dreams from controller
	dream, err := a.Controller.GetDream(id)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// write response
	sendResp(w, dream)
}

func (a API) GetInterpretation(w http.ResponseWriter, r *http.Request) {
	// get keyword from url
	vars := mux.Vars(r)
	keyword := vars["keyword"]

	// get interpretation from controller
	interpret, err := a.Controller.GetInterpret(keyword)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// write response
	sendResp(w, interpret)
}

func (a API) GetKeywords(w http.ResponseWriter, r *http.Request) {
	// get top from query
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "10"
	}

	top, err := strconv.Atoi(limit)
	if err != nil {
		sendError(w, http.StatusBadRequest, "limit must be an integer")
		return
	}

	// get keywords from controller
	keywords, err := a.Controller.GetKeywords(top)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// write response
	sendResp(w, keywords)
}

func (a API) CreateDream(w http.ResponseWriter, r *http.Request) {
	// get dream from body
	var dream models.Dream
	if err := json.NewDecoder(r.Body).Decode(&dream); err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	// create dream in controller
	id, err := a.Controller.WriteDreams(dream)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// write response
	sendResp(w, id)
}

func (a API) DeleteDream(w http.ResponseWriter, r *http.Request) {
	// get id from url
	vars := mux.Vars(r)
	id := vars["id"]

	// delete dream in controller
	if err := a.Controller.DeleteDream(id); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// write response
	sendResp(w, "ok")
}
