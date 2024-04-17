package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		RequiresAuth: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Function:     controllers.FindUsers,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodGet,
		Function:     controllers.FindUser,
		RequiresAuth: false,
	},
	{
		URI:          "/users/{userId}",
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
