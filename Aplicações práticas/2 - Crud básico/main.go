package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Code

	router := mux.NewRouter()

	fmt.Println("Server running on port 5000")

	log.Fatal(http.ListenAndServe(":5000", router))
}
