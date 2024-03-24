package main

import "fmt"

func main() {
	soma := somar(2, 7)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("erikinha")

	resultadoSoma, resultadoSubstracao := calculosMatematicos(2, 3)
	fmt.Println(resultadoSoma, resultadoSubstracao)

	resultado, _ := calculosMatematicos(2, 3)
	fmt.Println(resultado)
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao

}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}
