package models

type Interpretation struct {
	DreamID        int64  `json:"dream_id"`
	Title          string `json:"title"`
	Interpretation string `json:"interpretation"`
}
