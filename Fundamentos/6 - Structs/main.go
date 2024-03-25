package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     string
}

func main() {
	var u usuario
	u.nome = "erika"
	u.idade = 23
	fmt.Println(u)

	u2 := usuario{"erikinha", 23, endereco{"rua abc", "47"}}
	fmt.Println(u2)

	usuario3 := usuario{nome: "Erika"}
	fmt.Println(usuario3)
}
