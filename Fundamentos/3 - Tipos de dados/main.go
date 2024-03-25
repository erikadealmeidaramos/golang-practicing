package main

import (
	"errors"
	"fmt"
)

func main() {
	char := 'B'

	fmt.Println(char)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
