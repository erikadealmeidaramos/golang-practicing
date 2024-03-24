package main

import "fmt"

func main() {
	// Função anônima
	func() {
		fmt.Println("Função anônima")
	}()

	// Função anônima com parâmetros
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

	// Função anônima com retorno
	resultado := func(x, y int) int {
		return x + y
	}(10, 20)
	fmt.Println(resultado)

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parâmetro")

	fmt.Println(retorno)
}
