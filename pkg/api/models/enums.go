package models

// create an enum
type WebApp int

const (
	GorillaMux WebApp = iota
	Fiber
)
