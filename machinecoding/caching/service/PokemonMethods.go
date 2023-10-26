package service

import (
	cache "caching/cache"
	domain "caching/domain"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Get Pokemon data by Id
// Get Pokemon data by Name
// Delete Pokemon data by Id
// Add Pokemon data

var cacheObj *cache.LRUCaching

func Initialization(capacity int) {
	cacheObj = cache.InItCache(capacity)
}
func GetPokmonDetailsById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//fetching data from the request body by id
	var pokemon []domain.Pokemon
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

func AddPokemonDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//reading the request body from the controller function
	var pokemon domain.Pokemon
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
	var pokemon []domain.Pokemon
	pokemon = cacheObj.ListAll()
	json.NewEncoder(w).Encode(pokemon)
}

func DeletePokemonById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	cacheResponse := cacheObj.DeleteById(id)
	if cacheResponse.Err != nil {
		http.Error(w, cacheResponse.Err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(cacheResponse.Pokemon)
}
