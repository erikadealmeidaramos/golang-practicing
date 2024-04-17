package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI          string
	Method       string
	Function     func(w http.ResponseWriter, r *http.Request)
	RequiresAuth bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {
		if route.RequiresAuth {

			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Auth(route.Function))).Methods(route.Method)

		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}

	}

	return r
}
