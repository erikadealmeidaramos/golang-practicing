package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	userSaved, error := repository.FindByEmail(user.Email)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if error = security.VerifyPassword(userSaved.Password, user.Password); error != nil {
		responses.Error(w, http.StatusUnauthorized, fmt.Errorf("invalid password"))
		return
	}

	token, error := auth.CreateToken(userSaved.ID)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	w.Write([]byte(token))

}
