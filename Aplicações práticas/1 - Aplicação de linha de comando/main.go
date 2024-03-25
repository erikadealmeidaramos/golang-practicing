package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()

	/*
		Exemplos de como chamar:
		go run main.go ip --host mercadolivre.com.br
		go run main.go servidores --host mercadolivre.com.br
	*/

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
