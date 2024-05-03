package auth

import (
	"api/src/config"
	"errors"
	"net/http"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func CreateToken(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}

	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)

	token, error := jwt.Parse(tokenString, returnVerificationKey)

	if error != nil {
		return error
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, jwt.ErrSignatureInvalid
	}
	return config.SecretKey, nil
}

func ExtractUserId(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)

	token, error := jwt.Parse(tokenString, returnVerificationKey)

	if error != nil {
		return 0, error
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := permissions["userId"]
		if userId != nil {
			return uint64(userId.(float64)), nil
		}
	}

	return 0, errors.New("invalid token")
}
