package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	estudante1 := estudante{}
	estudante1.nome = "Erika"
	estudante1.sobrenome = "Ramos"
	estudante1.idade = 23
	estudante1.altura = 160
	estudante1.curso = "ADS"
	estudante1.faculdade = "Fatec"

	fmt.Println(estudante1)

	estudante2 := estudante{pessoa{"Erika", "Ramos", 23, 160}, "ADS", "Fatec"}

	fmt.Println(estudante2)

	estudante3 := estudante{
		pessoa: pessoa{
			nome:      "Erika",
			sobrenome: "Ramos",
			idade:     23,
			altura:    160,
		},
		curso:     "ADS",
		faculdade: "Fatec",
	}

	fmt.Println(estudante3)
}
