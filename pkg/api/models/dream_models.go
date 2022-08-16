package models

import "dreamt/pkg/models"

// create getDreamRequest struct
type GetDreamRequest struct {
	ID string `json:"id"`
}

type APIResponse struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
	Err    error
}

// create getDreamRequest struct
type GetInterpretationRequest struct {
	Keyword string `json:"keyword"`
}

// create getDreamRequest struct
type GetKeywordsRequest struct {
	Limit string `json:"limit"`
}

// create getDreamRequest struct
type CreateDreamRequest struct {
	Dream models.Dream `json:"dream"`
}

// create getDreamRequest struct
type DeleteDreamRequest struct {
	ID string `json:"id"`
}
