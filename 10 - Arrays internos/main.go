package main

import (
	"fmt"
)

func main() {

	/*a função make aloca espaço na memória para determinada coisa.
	Inicialmente o make cria um array e não um slice, de 15 posições (capacidade máxima) e
	depois cria um slice de 10 posições (tamanho inicial) que aponta para o array criado.
	*/

	slice := make([]float32, 10, 15)

	/*
		Se você estourar a capacidade máxima do slice, o Go cria um novo array com o
		dobro da capacidade do array anterior e copia os elementos do array anterior para o novo array.
	*/

	slice = append(slice, 4)
	slice = append(slice, 10)
	slice = append(slice, 12)
	slice = append(slice, 14)
	slice = append(slice, 16)
	slice = append(slice, 18)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
