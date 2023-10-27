package tests

import (
	ca "caching/cache"
	domain "caching/domain"
	"testing"
)

func TestLRUCaching_Get(t *testing.T) {
	cache := ca.InItCache(5)
	pokemon := domain.Pokemon{Id: 1, Name: "Pikachu"}

	// Test case 1: Key exists
	cache.Put(1, pokemon)
	result := cache.Get(1)

	if result.Err != nil || result.Pokemon != pokemon {
		t.Error("Get did not work as expected for an existing key.")
	}

	// Test case 2: Key doesn't exist
	result = cache.Get(2)

	if result.Err == nil || result.Pokemon != (domain.Pokemon{}) {
		t.Error("Get did not work as expected for a non-existing key.")
	}
}

func TestLRUCaching_Put(t *testing.T) {
	cache := ca.InItCache(2)
	pokemon1 := domain.Pokemon{Id: 1, Name: "Pikachu"}
	pokemon2 := domain.Pokemon{Id: 2, Name: "Charmander"}

	// Test case 1: Add elements within capacity
	cache.Put(1, pokemon1)
	cache.Put(2, pokemon2)

	if cache.Dll.Len != 2 {
		t.Error("Put did not work as expected within capacity.")
	}

	// Test case 2: Evict the oldest element
	pokemon3 := domain.Pokemon{Id: 3, Name: "Bulbasaur"}
	cache.Put(3, pokemon3)

	if cache.Dll.Len != 2 {
		t.Error("Put did not work as expected when evicting the oldest element.")
	}

	// Test case 3: Update an existing element
	pokemon1Updated := domain.Pokemon{Id: 1, Name: "Updated Pikachu"}
	cache.Put(1, pokemon1Updated)

	result := cache.Get(1)
	if result.Pokemon != pokemon1Updated {
		t.Error("Put did not update an existing element as expected.")
	}
}

func TestLRUCaching_DeleteById(t *testing.T) {
	cache := ca.InItCache(5)
	pokemon1 := domain.Pokemon{Id: 1, Name: "Pikachu"}
	//pokemon2 := domain.Pokemon{Id: 2, Name: "Charmander"}

	// Test case 1: Delete an existing element
	cache.Put(1, pokemon1)
	result := cache.DeleteById(1)

	if result.Err != nil || result.Pokemon != pokemon1 {
		t.Error("DeleteById did not work as expected for an existing key.")
	}

	// Test case 2: Delete a non-existing element
	result = cache.DeleteById(2)

	if result.Err == nil || result.Pokemon != (domain.Pokemon{}) {
		t.Error("DeleteById did not work as expected for a non-existing key.")
	}
}
