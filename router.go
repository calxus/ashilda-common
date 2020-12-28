package models

import "github.com/gorilla/mux"

// Router type for holding mux router and database connection used in handlers
type Router struct {
	basePath string
	database Database
	mux      *mux.Router
}

// NewRouter method to construct Router type
func NewRouter(bp string) *Router {
	return &Router{
		basePath: bp,
		database: *NewDatabase(),
		mux:      mux.NewRouter().StrictSlash(true),
	}
}
