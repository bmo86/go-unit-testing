package main

import (
	"go-testing-pokeAPI/controller"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/poke/{id}", controller.GetPokemon).Methods(http.MethodGet)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Printf("Found Server :c")
	}

	log.Printf("Server mount in port : ", 8080)

}
