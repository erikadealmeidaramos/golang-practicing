package main

import (
	"fmt"
)

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func main() {
	fmt.Println("Dia 1: ", diaDaSemana(1))
	fmt.Println("Dia 2: ", diaDaSemana(2))
	fmt.Println("Dia 3: ", diaDaSemana(3))
	fmt.Println("Dia 4: ", diaDaSemana(4))
	fmt.Println("Dia 5: ", diaDaSemana(5))
	fmt.Println("Dia 6: ", diaDaSemana(6))
	fmt.Println("Dia 7: ", diaDaSemana(7))
	fmt.Println("Dia 8: ", diaDaSemana(8))
}
