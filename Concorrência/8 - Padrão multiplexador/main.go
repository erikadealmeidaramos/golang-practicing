package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá Mundo"), escrever("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

/*Recebe dois canais como entrada e retorna apenas um como saída*/
func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			/*
				Esse select vai verificar qual canal tem uma mensagem para ser lida,
				mas indendente de qual canal seja, ele fai jogar para o canal de saída
			*/
			case mensagem := <-canalDeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <-canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}

	}()

	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canal
}
