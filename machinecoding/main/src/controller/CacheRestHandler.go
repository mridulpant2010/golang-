package main

import (
	"fmt"
	"log"
	"net/http"

	service "pokemon.com/cache/src/service"

	"github.com/gorilla/mux"
)

// Get Pokemon data by Id
// Get Pokemon data by Name
// Delete Pokemon data by Id
// Add Pokemon data

func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/pokemon/{id}", service.GetPokmonDetailsById).Methods("GET")
	r.HandleFunc("/api/v1/pokemons/", service.AddPokemonDetails).Methods("POST")
	r.HandleFunc("/api/v1/pokemon/{id}", service.DeletePokemonById).Methods("DELETE")
	r.HandleFunc("/api/v1/pokemons", service.ListAllPokemon).Methods("GET")
	port := 8080
	fmt.Printf("Server is running on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

func main() {

	cacheSize := 5
	service.Initialization(cacheSize)
	initializeRouter()
}
