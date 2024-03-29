package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	/*
		Se você não quiser que algum campo seja serializado, basta colocar um traço
		no lugar do nome do campo. Exemplo: `json:"-"`
	*/
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJSON := `{"nome":"Rex","raca":"Dálmata","idade":3}`

	var c cachorro

	/*
		1. Parâmetro: os dados que a gente quer passar. Precisa ser um slice de bytes,
		por isso, utilizamos []byte para converter a string.
		2. Parâmetro: o endereço de memória da variável onde iremos armazenar os dados
	*/
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	c2 := make(map[string]string)

	cachorro2EmJSON := `{"nome":"Toby","raca":"Poodle"}`

	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
