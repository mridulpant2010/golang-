package algorithms_test

import (
	"testing"

	algo "pokemon.com/cache/src/algorithms"
	entity "pokemon.com/cache/src/entity"
)

func TestDoublyLinkedList_AddElementToEnd(t *testing.T) {
	dll := algo.NewDoublyLinkedList()
	pokemon := entity.Pokemon{Id: 1, Name: "Pikachu"}

	node := dll.AddElementsToEnd(1, pokemon)
	if dll.Len == 0 || dll.Head != node {
		t.Error("AddElementToEnd did not worked as expected")
	} else {
		t.Logf("PASSED: Expected DLL len %d, got %d", 1, dll.Len)
	}

}

func TestDoublyLinkedList_RemoveFromBeginning(t *testing.T) {
	dll := algo.NewDoublyLinkedList()
	pokemon := entity.Pokemon{Id: 1, Name: "Pikachu"}
	dll.AddElementsToEnd(1, pokemon)

	dll.RemoveFromBeginning()
	if dll.Len != 0 || dll.Head != nil || dll.Tail != nil {
		t.Error("RemoveFromBeginning did not worked as expected")
	} else {
		t.Logf("PASSED: Expected DLL len %d, got %d", 0, dll.Len)
	}

}

func TestDoublyLinkedList_MoveToEnd(t *testing.T) {
	dll := algo.NewDoublyLinkedList()
	pokemon := entity.Pokemon{Id: 1, Name: "Pikachu"}
	pokemon2 := entity.Pokemon{Id: 2, Name: "Charmender"}
	node1 := dll.AddElementsToEnd(1, pokemon)
	node2 := dll.AddElementsToEnd(2, pokemon2)
	dll.MoveToEnd(node1)

	if node2.Next != node1 || dll.Head != node2 || dll.Tail != node1 {
		t.Error("MoveToEnd did not worked as expected")
	}
}

func TestDoublyLinkedList_ListElements(t *testing.T) {
	dll := algo.NewDoublyLinkedList()
	pokemon := entity.Pokemon{Id: 1, Name: "Pikachu"}
	pokemon2 := entity.Pokemon{Id: 2, Name: "Charmender"}
	dll.AddElementsToEnd(1, pokemon)
	dll.AddElementsToEnd(2, pokemon2)

	pokemonList := dll.ListElements()
	if len(pokemonList) != 2 {
		t.Error("TestDoublyLinkedList_ListElements did not worked as expected")
	} else {
		t.Logf("PASSED: Expected DLL len %d, got %d", 2, dll.Len)
	}

}
