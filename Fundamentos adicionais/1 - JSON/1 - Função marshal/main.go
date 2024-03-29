package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	/*
		Aqui abaixo ele criou um slice de bytes unit8
	*/
	fmt.Println(cachorroEmJSON)
	/*
		Para visualizar o JSON, é necessário converter o slice de bytes para string.
		Para isso, utilizamos o pacote bytes e a função NewBuffer.
	*/
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))
}
