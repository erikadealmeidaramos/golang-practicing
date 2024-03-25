package main

import "fmt"

func main() {
	var var1 string = "Variável 1"
	fmt.Println(var1)

	var2 := "Variável 2"
	fmt.Println(var2)

	var (
		variavel3 string = "lala"
		variavel4 string = "lala2"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "lala3", "lala4"

	fmt.Println(variavel5, variavel6)

	variavel5, variavel6 = variavel6, variavel5
}
