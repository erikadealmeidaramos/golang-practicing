package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"

	"api/src/config"
)

func main() {
	config.Load()

	fmt.Println("Escutando na porta: ", fmt.Sprintf(":%d", config.Port))

	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
