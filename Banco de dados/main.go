package main

import (
	"database/sql"
	"fmt"
	"log"

	/*
		O underline é utilizado para importar um pacote sem utilizá-lo diretamente no código.
		Nesse caso, é necessário pois o pacote é utilizado apenas para registrar o driver do MySQL.
	*/
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)

	/*
		Esse trecho detecta erros mais gerais (fora autenticação)
	*/
	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	/*
		Esse trecho detecta erros de autenticação, checando se a conexão está ativa
		com o ping.
	*/
	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	/*
		Linhas é uma espécie de cursor que percorre o resultado da query.
		Para ver todas as linhas, é necessário percorrer o cursor com um looping.
	*/
	linhas, erro := db.Query("select * from usuarios")

	if erro != nil {
		log.Fatal(erro)
	}

	defer linhas.Close()

	fmt.Println(linhas)
}
