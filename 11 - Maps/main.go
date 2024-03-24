package main

import "fmt"

func main() {
	//dentro do colchete é o tipo da chave e fora o tipo do valor
	usuario := map[string]string{
		"nome":      "Erika",
		"sobrenome": "Ramos",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Erika",
			"ultimo":   "Ramos",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["nome"] = map[string]string{
		"primeiro": "Erika",
		"ultimo":   "Ramos",
	}
	fmt.Println(usuario2)
}
