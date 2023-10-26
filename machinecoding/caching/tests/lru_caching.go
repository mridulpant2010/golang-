package main

import (
	ca "caching/cache"
	domain "caching/domain"
)

func main() {
	cacheObj := ca.InItCache(5)
	//fmt.Println(cacheObj.Capacity, cacheObj.Dll, cacheObj.Storage)
	pokemonDataset := []domain.Pokemon{
		{Id: 1, Name: "pikachu", Height: 3.4, Weight: 40.3, Type: "htype", Abilities: "can swim"},
		{Id: 2, Name: "bulbasaur", Height: 5.4, Weight: 60.2, Type: "ftype", Abilities: "can fire"},
		{Id: 3, Name: "raichu", Height: 4.3, Weight: 32, Type: "gtype", Abilities: "can bullet"},
	}

	for k, v := range pokemonDataset {
		cacheObj.Put(k+1, v)
	}

	// cacheObj.Put(1, "one")
	// cacheObj.Put(2, "two")
	// cacheObj.Put(3, "three")
	// cacheObj.Put(4, "four")
	// cacheObj.Put(5, "five")
	// cacheObj.Put(6, "six")
	cacheObj.ListAll()

	cacheObj.Dll.Display()
	cacheObj.Get(2)
	cacheObj.Dll.Display()
	cacheObj.Put(3, domain.Pokemon{Id: 3, Name: "niniachu", Height: 3.4, Weight: 40.3, Type: "htype", Abilities: "can swim"})
	cacheObj.Dll.Display()
	cacheObj.Get(1)
}
