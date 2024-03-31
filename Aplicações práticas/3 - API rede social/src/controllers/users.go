package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando usuários"))
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando usuário"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}
