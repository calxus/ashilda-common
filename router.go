package models

import "github.com/gorilla/mux"

// Router type for holding mux router and database connection used in handlers
type Router struct {
	BasePath string
	Database Database
	Mux      *mux.Router
}

// NewRouter method to construct Router type
func NewRouter(bp string) *Router {
	return &Router{
		BasePath: bp,
		Database: *NewDatabase(),
		Mux:      mux.NewRouter().StrictSlash(true),
	}
}
