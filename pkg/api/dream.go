package api

import (
	rmodels "dreamt/pkg/api/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (a API) getDreams() rmodels.APIResponse {
	// get dreams from controller
	dreams, err := a.controller.GetDreams()
	if err != nil {
		fmt.Println("err: ", err)
		return rmodels.APIResponse{
			Status: http.StatusInternalServerError,
			Err:    err,
		}
	}

	// write response
	return rmodels.APIResponse{Status: http.StatusOK, Body: dreams}
}

func (a API) getDream(r rmodels.GetDreamRequest) rmodels.APIResponse {
	// get dreams from controller
	dream, err := a.controller.GetDream(r.ID)
	if err != nil {
		fmt.Println("err: ", err)
		return rmodels.APIResponse{
			Status: http.StatusInternalServerError,
			Err:    err,
		}
	}

	// write response
	return rmodels.APIResponse{Body: dream, Status: http.StatusOK}
}

func (a API) getInterpret(r rmodels.GetInterpretationRequest) rmodels.APIResponse {
	// get interpretation from controller
	interpret, err := a.controller.GetInterpret(r.Keyword)
	if err != nil {
		return rmodels.APIResponse{
			Status: http.StatusInternalServerError,
			Err:    err,
		}
	}

	// write response
	return rmodels.APIResponse{Body: interpret, Status: http.StatusOK}
}

func (a API) getKeywords(r rmodels.GetKeywordsRequest) rmodels.APIResponse {
	// get top from query
	limit := r.Limit
	if limit == "" {
		limit = "10"
	}

	top, err := strconv.Atoi(limit)
	if err != nil {
		return rmodels.APIResponse{
			Status: http.StatusBadRequest,
			Err:    errors.New("limit must be an integer: " + err.Error()),
		}
	}

	// get keywords from controller
	keywords, err := a.controller.GetKeywords(top)
	if err != nil {
		return rmodels.APIResponse{
			Status: http.StatusInternalServerError,
			Err:    err,
		}
	}

	// write response
	return rmodels.APIResponse{Status: http.StatusOK, Body: keywords}
}

func (a API) createDream(r rmodels.CreateDreamRequest) rmodels.APIResponse {
	// create dream in controller
	id, err := a.controller.WriteDreams(r.Dream)
	if err != nil {
		fmt.Println("err: ", err)
		return rmodels.APIResponse{
			Status: http.StatusInternalServerError,
			Err:    err,
		}
	}

	// write response
	return rmodels.APIResponse{Status: http.StatusCreated, Body: id}
}

func (a API) deleteDream(r rmodels.DeleteDreamRequest) rmodels.APIResponse {
	// get id from url
	// delete dream in controller
	if err := a.controller.DeleteDream(r.ID); err != nil {
		return rmodels.APIResponse{
			Status: http.StatusInternalServerError,
			Err:    err,
		}
	}

	// write response
	return rmodels.APIResponse{Status: http.StatusOK, Body: "ok"}
}
