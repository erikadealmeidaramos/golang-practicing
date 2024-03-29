package main

import "fmt"

func fibonacci(x int) int {
	if x <= 1 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

/*
O <- antes do chan (que indica o canal) representa que o canal apenas recebe
dados, ou seja, é um canal somente leitura.

Para indicar que o canal só envia dados, o símbolo que representa é o <- depois
do chan. Ou seja, é um canal somente escrita.
*/
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}
