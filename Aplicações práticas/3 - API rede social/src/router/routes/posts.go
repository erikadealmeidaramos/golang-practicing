package routes

import (
	"api/src/controllers"
	"net/http"
)

var postRoutes = []Route{
	{
		URI:          "/posts",
		Method:       http.MethodPost,
		Function:     controllers.CreatePost,
		RequiresAuth: true,
	},
	{
		URI:          "/posts",
		Method:       http.MethodGet,
		Function:     controllers.FindPosts,
		RequiresAuth: true,
	},
	{
		URI:          "/posts/{postId}",
		Method:       http.MethodGet,
		Function:     controllers.FindPost,
		RequiresAuth: true,
	},
	{
		URI:          "/posts/{postId}",
		Method:       http.MethodPut,
		Function:     controllers.UpdatePost,
		RequiresAuth: true,
	},
	{
		URI:          "/posts/{postId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeletePost,
		RequiresAuth: true,
	},
	{
		URI:          "/posts/{postId}/like",
		Method:       http.MethodPost,
		Function:     controllers.LikePost,
		RequiresAuth: true,
	},
	{
		URI:          "/posts/{postId}/unlike",
		Method:       http.MethodDelete,
		Function:     controllers.UnlikePost,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userId}/posts",
		Method:       http.MethodGet,
		Function:     controllers.FindPostsByUser,
		RequiresAuth: true,
	},
}
