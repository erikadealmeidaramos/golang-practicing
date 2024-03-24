package main

import "fmt"

func fatorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * fatorial(x-1)
}

func fibonacci(x uint) uint {
	if x <= 1 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

func main() {
	// Função recursiva
	fmt.Println(fatorial(5))
	fmt.Println(fibonacci(10))

	posicao := uint(12)

	//se eu quiser imprimir todos os números da sequência
	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
