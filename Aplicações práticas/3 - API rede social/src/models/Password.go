package models

type Password struct {
	New     string `json:"new" validate:"required,min=6"`
	Current string `json:"current" validate:"required,min=6"`
}
