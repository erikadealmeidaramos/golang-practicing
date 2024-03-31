package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Function:     func(w http.ResponseWriter, r *http.Request) {},
		RequiresAuth: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Function:     controllers.FindUsers,
		RequiresAuth: false,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodGet,
		Function:     controllers.FindUser,
		RequiresAuth: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		RequiresAuth: false,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		RequiresAuth: false,
	},
}
