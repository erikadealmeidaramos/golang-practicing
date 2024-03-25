package main

import "fmt"

func main() {
	var numero int8 = 10
	numero++
	fmt.Println(numero)

	if numero == 8 {
		fmt.Println("É igual a 8")
	} else {
		fmt.Println("Não é igual a 8")
	}
}
