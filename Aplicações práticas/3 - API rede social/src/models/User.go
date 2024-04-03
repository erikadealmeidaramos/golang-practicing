package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	/*O omitempty omite o campo do json caso o valor esteja vazio*/
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("o campo nome é obrigatório e não pode estar em branco")
	}
	if user.Nick == "" {
		return errors.New("o campo nick é obrigatório e não pode estar em branco")
	}
	if user.Email == "" {
		return errors.New("o campo email é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("o campo email está inválido")
	}

	if step == "register" && user.Password == "" {
		return errors.New("o campo senha é obrigatório e não pode estar em branco")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}

func (user *User) Prepare(step string) error {
	if error := user.validate(step); error != nil {
		return error
	}

	user.format()

	return nil
}
