package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá Mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	/*Aqui estou recebendo o valor que foi enviado para o canal.
	Ele vai aguardar até que o valor seja enviadao para o canal, antes
	de prosseguir com a execução do código.
	*/
	/*for {
		mensagem, aberto := <-canal

		if !aberto {
			break
		}

		fmt.Println(mensagem)
	}*/

	/*
		Outra maneira de fazer o código acima:
	*/

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	//time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		//Aqui estou mandando um valor para dentro do canal
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
