package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel int = 10
	var variavel2 int = variavel

	variavel++
	fmt.Println(variavel, variavel2)

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)

	fmt.Println(*ponteiro)

}
