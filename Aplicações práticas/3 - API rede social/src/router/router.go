package router

import "github.com/gorilla/mux"

// Generate creates a new router with the default configuration
func Generate() *mux.Router {
	return mux.NewRouter()
}
