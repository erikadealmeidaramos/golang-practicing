package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userId, error := auth.ExtractUserId(r)

	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var post models.Post
	error = json.Unmarshal(requestBody, &post)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	post.AuthorID = userId

	if error := post.Prepare(); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewPostRepository(db)
	post.ID, error = repository.Create(post)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusCreated, post)

}

func FindPosts(w http.ResponseWriter, r *http.Request) {
	userId, error := auth.ExtractUserId(r)

	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewPostRepository(db)
	posts, error := repository.FindAll(userId)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

func FindPost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postId, error := strconv.ParseUint(parameters["postId"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewPostRepository(db)
	posts, error := repository.FindById(postId)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userId, error := auth.ExtractUserId(r)

	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	parameters := mux.Vars(r)
	postId, error := strconv.ParseUint(parameters["postId"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var post models.Post
	error = json.Unmarshal(requestBody, &post)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	post.AuthorID = userId
	post.ID = postId

	if error := post.Prepare(); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewPostRepository(db)
	postDb, error := repository.FindById(postId)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if postDb.AuthorID != userId {
		responses.Error(w, http.StatusForbidden, errors.New("you can only update your own posts"))
		return
	}

	if error := repository.Update(postId, post); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	userId, error := auth.ExtractUserId(r)

	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	parameters := mux.Vars(r)
	postId, error := strconv.ParseUint(parameters["postId"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewPostRepository(db)
	post, error := repository.FindById(postId)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if post.AuthorID != userId {
		responses.Error(w, http.StatusForbidden, errors.New("you can only delete your own posts"))
		return
	}

	if error := repository.Delete(postId); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, error := strconv.ParseUint(parameters["postId"], 10, 64)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewPostRepository(db)
	if error := repository.Like(postId); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func UnlikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, error := strconv.ParseUint(parameters["postId"], 10, 64)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewPostRepository(db)
	if error := repository.Unlike(postId); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func FindPostsByUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, error := strconv.ParseUint(parameters["userId"], 10, 64)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewPostRepository(db)
	posts, error := repository.FindPostByUserId(userId)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}
