package main

import (
	"log"
	"net/http"
)

/* HTTP é um protocolo de comunicação que permite a transferência de arquivos de texto,
imagem, vídeo, áudio, etc. na web.

CLIENTE (FAZ A REQUISIÇÃO) -> SERVIDOR (PROCESSA A REQUISIÇÃO) -> CLIENTE (RECEBE A RESPOSTA)

REQUEST (REQUISIÇÃO) -> RESPONSE (RESPOSTA)

ROTAS (ENDPOINTS) -> URLS QUE O SERVIDOR RESPONDE
 // URI (IDENTIFICADOR DE RECURSO UNIVERSAL) -> você fala para o servidor o que você quer fazer
 (isso tem a ver com produtos, usuários, etc)

MÉTODOS HTTP: GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD
*/

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá, mundo!"))
}

func main() {

	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
