package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// criarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario Usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")

	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)

	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o ID inserido"))
		return
	}

	//STATUS CODES

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))

}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")

	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários"))
		return
	}

	defer linhas.Close()

	var usuarios []Usuario
	for linhas.Next() {
		var usuario Usuario
		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro = json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	defer db.Close()

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuário"))
		return
	}

	defer linha.Close()

	var usuario Usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
	}

	if usuario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuário não encontrado"))
		return
	}

	w.WriteHeader(http.StatusOK)

	if erro = json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para JSON"))
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	var usuario Usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")

	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")

	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
