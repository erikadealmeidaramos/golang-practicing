package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
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
