package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados...\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	fmt.Printf("Vendo se o usuário %s é maior de idade:\n", u.nome)
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	fmt.Printf("Adicionando um ano de idade ao usuário %s...\n", u.nome)
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Erikinha", 20}
	maiorIdade := usuario2.maiorDeIdade()

	fmt.Println(maiorIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)

}
