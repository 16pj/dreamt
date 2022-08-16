package models

type Dream struct {
	Id      int64    `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type DreamHeader struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}
