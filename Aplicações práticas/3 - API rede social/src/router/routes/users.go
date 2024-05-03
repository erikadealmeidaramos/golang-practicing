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
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userId}/follow",
		Method:       http.MethodPost,
		Function:     controllers.FollowUser,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userId}/unfollow",
		Method:       http.MethodDelete,
		Function:     controllers.UnfollowUser,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userId}/followers",
		Method:       http.MethodGet,
		Function:     controllers.FindFollowers,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userId}/following",
		Method:       http.MethodGet,
		Function:     controllers.FindFollowing,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userId}/update-password",
		Method:       http.MethodPatch,
		Function:     controllers.UpdatePassword,
		RequiresAuth: true,
	},
}
