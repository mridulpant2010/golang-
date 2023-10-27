package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gorilla/mux"
// 	entity "pokemon.com/cache/src/entity"
// 	service "pokemon.com/cache/src/service"
// )

// func TestGetPokmonDetailsById(t *testing.T) {
// 	// Initialize your cache and router
// 	service.Initialization(5)
// 	r := mux.NewRouter()
// 	r.HandleFunc("/api/v1/pokemon/{id}", service.GetPokmonDetailsById).Methods("GET")

// 	// Create a request with a valid ID
// 	req, _ := http.NewRequest("GET", "/api/v1/pokemon/1", nil)
// 	rr := httptest.NewRecorder()
// 	r.ServeHTTP(rr, req)
// 	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 8082), r))
// 	// Check the response status code
// 	if rr.Code != http.StatusOK {
// 		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
// 	}

// 	// Create a request with an invalid ID
// 	req, _ = http.NewRequest("GET", "/api/v1/pokemon/100", nil)
// 	rr = httptest.NewRecorder()
// 	r.ServeHTTP(rr, req)

// 	// Check the response status code for not found
// 	if rr.Code != http.StatusNotFound {
// 		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, rr.Code)
// 	}
// }

// func TestAddPokemonDetails(t *testing.T) {
// 	// Initialize your cache and router
// 	service.Initialization(5)
// 	r := mux.NewRouter()
// 	r.HandleFunc("/api/v1/pokemons", service.AddPokemonDetails).Methods("POST")

// 	// Create a Pokemon entity to add
// 	pokemon := entity.Pokemon{Id: 1, Name: "Pikachu"}

// 	// Marshal the Pokemon entity to JSON
// 	pokemonJSON, _ := json.Marshal(pokemon)

// 	// Create a request to add the Pokemon
// 	req, _ := http.NewRequest("POST", "/api/v1/pokemons", bytes.NewReader(pokemonJSON))
// 	req.Header.Set("Content-Type", "application/json")
// 	rr := httptest.NewRecorder()
// 	r.ServeHTTP(rr, req)

// 	// Check the response status code
// 	if rr.Code != http.StatusCreated {
// 		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rr.Code)
// 	}
// }
