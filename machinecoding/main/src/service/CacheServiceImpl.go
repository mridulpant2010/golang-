package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	ca "pokemon.com/cache/src/cache"
	entity "pokemon.com/cache/src/entity"

	"github.com/gorilla/mux"
)

// Get Pokemon data by Id
// Delete Pokemon data by Id
// Add Pokemon data

var cacheObj ca.Caching

func Initialization(capacity int) {
	cacheObj = ca.InItCache(capacity) //initialize a LRU cache
}

// INFO: function to fetch the pokemon details by Id
func GetPokmonDetailsById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//fetching data from the request body by id
	var pokemon []entity.Pokemon
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	cacheResponse := cacheObj.Get(id)
	pokemon = append(pokemon, cacheResponse.Pokemon)
	if cacheResponse.Err != nil {
		http.Error(w, cacheResponse.Err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(pokemon)
}

// INFO: function to add the pokemon details
func AddPokemonDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//reading the request body from the controller function
	var pokemon entity.Pokemon
	err := json.NewDecoder(r.Body).Decode(&pokemon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cacheObj.Put(pokemon.Id, pokemon)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pokemon)
}

func ListAllPokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var pokemon []entity.Pokemon
	pokemon = cacheObj.ListAll()
	json.NewEncoder(w).Encode(pokemon)
}

func DeletePokemonById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	cacheResponse := cacheObj.DeleteById(id)
	//return NotFound if the key does not exist
	if cacheResponse.Err != nil {
		http.Error(w, cacheResponse.Err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("Element deleted")
}
