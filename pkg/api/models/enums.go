package models

// create an enum
type WebApp string

const (
	GorillaMux WebApp = "gorilla"
	Fiber      WebApp = "fiber"
)
