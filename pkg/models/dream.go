package models

type Dream struct {
	ID      int64    `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type DreamHeader struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}
