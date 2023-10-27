package cache_test

import (
	"fmt"
	"testing"

	ca "pokemon.com/cache/src/cache"
	entity "pokemon.com/cache/src/entity"
)

func TestLRUCaching_Get(t *testing.T) {
	cache := ca.InItCache(5)
	pokemon := entity.Pokemon{Id: 1, Name: "Pikachu"}

	cache.Put(1, pokemon)
	result := cache.Get(1)
	// case1: key exists
	if result.Err != nil || result.Pokemon != pokemon {
		t.Error("Get did not worked as expected for an existing key")
	}

	//case 2: key does not exist
	result = cache.Get(2)
	err_msg := fmt.Errorf("Key doesnt exist")
	if result.Err == nil {
		t.Error("Get did not worked as expected for an non-existing key")
	} else {
		t.Logf("PASSED: Expected error %s, got %s", err_msg, result.Err)
	}

}

func TestLRUCaching_Put(t *testing.T) {
	cache := ca.InItCache(2)
	pokemon1 := entity.Pokemon{Id: 1, Name: "Pikachu"}
	pokemon2 := entity.Pokemon{Id: 2, Name: "Charmander"}

	// case 1: Add elements within capacity
	cache.Put(1, pokemon1)
	cache.Put(2, pokemon2)

	if cache.Length() != 2 {
		t.Error("Put did not work as expected within capacity.")
	} else {
		t.Logf("PASSED case1: Expected Cache len %d, got %d", 2, cache.Length())
	}

	// Test case 2: Evict the oldest element
	pokemon3 := entity.Pokemon{Id: 3, Name: "Bulbasaur"}
	cache.Put(3, pokemon3)

	if cache.Length() != 2 {
		t.Error("Put did not work as expected within capacity.")
	} else {
		t.Logf("PASSED case2: Expected Cache len %d, got %d", 2, cache.Length())
	}

	// Test case 3: Update an existing element
	pokemon1Updated := entity.Pokemon{Id: 1, Name: "Updated Pikachu"}
	cache.Put(1, pokemon1Updated)

	result := cache.Get(1)
	if result.Pokemon != pokemon1Updated {
		t.Error("Put did not update an existing element as expected.")
	}
}

func TestLRUCaching_DeleteById(t *testing.T) {
	cache := ca.InItCache(5)
	pokemon1 := entity.Pokemon{Id: 1, Name: "Pikachu"}
	//pokemon2 := domain.Pokemon{Id: 2, Name: "Charmander"}

	// Test case 1: Delete an existing element
	cache.Put(1, pokemon1)
	result := cache.DeleteById(1)

	if result.Err != nil || result.Pokemon != pokemon1 {
		t.Error("DeleteById did not work as expected for an existing key.")
	}

	// Test case 2: Delete a non-existing element
	result = cache.DeleteById(2)

	if result.Err == nil || result.Pokemon != (entity.Pokemon{}) {
		t.Error("DeleteById did not work as expected for a non-existing key.")
	}
}
