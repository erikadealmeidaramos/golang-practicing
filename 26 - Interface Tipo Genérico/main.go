package main

import (
	"fmt"
)

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(1.5)
	generica(true)

	//só é bom usar em casos muito específicos, como na função de imprimir coisa na tela, por exemplo, que realmente faz sentido nesse contexto
	fmt.Println(1, "String", false, true, float32(5688))
}
