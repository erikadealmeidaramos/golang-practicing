package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, error := io.ReadAll(r.Body)

	if error != nil {
		log.Fatal(error)
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.Prepare("register"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	user.ID, error = repository.Create(user)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusCreated, user)

}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, error := database.Connect()

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	users, error := repository.Find(nameOrNick)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
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

	repository := repositories.NewUsersRepository(db)

	user, error := repository.FindByID(userId)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, error := strconv.ParseUint(parameters["userId"], 10, 64)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	tokenUserId, error := auth.ExtractUserId(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return

	}

	if userId != tokenUserId {
		responses.Error(w, http.StatusForbidden, fmt.Errorf("you can only update your own profile"))
	}

	requestBody, error := io.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.Prepare("update"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	if error = repository.Update(userId, user); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, error := strconv.ParseUint(parameters["userId"], 10, 64)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	tokenUserId, error := auth.ExtractUserId(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return

	}

	if userId != tokenUserId {
		responses.Error(w, http.StatusForbidden, fmt.Errorf("you can only delete your own profile"))
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	if error = repository.Delete(userId); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, error := auth.ExtractUserId(r)

	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	parameters := mux.Vars(r)
	userId, error := strconv.ParseUint(parameters["userId"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if followerId == userId {
		responses.Error(w, http.StatusForbidden, fmt.Errorf("you can't follow yourself"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	if error = repository.Follow(userId, followerId); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, error := auth.ExtractUserId(r)

	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	parameters := mux.Vars(r)
	userId, error := strconv.ParseUint(parameters["userId"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if followerId == userId {
		responses.Error(w, http.StatusForbidden, fmt.Errorf("you can't unfollow yourself"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	if error = repository.Unfollow(userId, followerId); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func FindFollowers(w http.ResponseWriter, r *http.Request) {
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

	repository := repositories.NewUsersRepository(db)
	followers, error := repository.FindFollowers(userId)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

func FindFollowing(w http.ResponseWriter, r *http.Request) {
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

	repository := repositories.NewUsersRepository(db)
	following, error := repository.FindFollowing(userId)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, following)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userId, error := auth.ExtractUserId(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	parameters := mux.Vars(r)
	userIdPath, error := strconv.ParseUint(parameters["userId"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if userId != userIdPath {
		responses.Error(w, http.StatusForbidden, fmt.Errorf("you can only update your own password"))
		return
	}

	requestBody, error := io.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var password models.Password
	if error = json.Unmarshal(requestBody, &password); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	passwordDb, error := repository.FindPassword(userId)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if error = security.VerifyPassword(passwordDb, password.Current); error != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("the password is incorrect"))
		return
	}

	hashPassword, error := security.Hash(password.New)

	password.New = string(hashPassword)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = repository.UpdatePassword(userId, password.New); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
