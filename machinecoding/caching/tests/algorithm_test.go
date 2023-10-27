package tests

import (
	algo "caching/algorithms"
	"caching/domain"
	"testing"
)

func TestDoublyLinkedList_AddElementToEnd(t *testing.T) {
	dll := algo.NewDoublyLinkedList()
	pokemon := domain.Pokemon{Id: 1, Name: "Pikachu"}

	node := dll.AddElementToEnd(1, pokemon)

	if dll.Head != node || dll.Tail != node || dll.Len != 1 {
		t.Error("AddElementToEnd did not work as expected.")
	}
}

func TestDoublyLinkedList_RemoveFromBeginning(t *testing.T) {
	dll := algo.NewDoublyLinkedList()
	pokemon := domain.Pokemon{Id: 1, Name: "Pikachu"}
	dll.AddElementToEnd(1, pokemon)

	node := dll.RemoveFromBeginning()

	if node == nil || node.Data.Value != pokemon || dll.Len != 0 || dll.Head != nil || dll.Tail != nil {
		t.Error("RemoveFromBeginning did not work as expected.")
	}
}

func TestDoublyLinkedList_MoveToEnd(t *testing.T) {
	dll := algo.NewDoublyLinkedList()
	pokemon1 := domain.Pokemon{Id: 1, Name: "Pikachu"}
	pokemon2 := domain.Pokemon{Id: 2, Name: "Charmander"}

	node1 := dll.AddElementToEnd(1, pokemon1)
	node2 := dll.AddElementToEnd(2, pokemon2)

	dll.MoveToEnd(node1)

	if dll.Head != node2 || dll.Tail != node1 || node2.Next != node1 || node1.Prev != node2 {
		t.Error("MoveToEnd did not work as expected.")
	}
}

func TestDoublyLinkedList_DetachNode(t *testing.T) {
	dll := algo.NewDoublyLinkedList()
	pokemon1 := domain.Pokemon{Id: 1, Name: "Pikachu"}
	pokemon2 := domain.Pokemon{Id: 2, Name: "Charmander"}

	node1 := dll.AddElementToEnd(1, pokemon1)
	node2 := dll.AddElementToEnd(2, pokemon2)

	dll.DetachNode(node1)

	if dll.Head != node2 || dll.Tail != node2 || node2.Next != nil || node1.Prev != nil {
		t.Error("DetachNode did not work as expected.")
	}
}

func TestDoublyLinkedList_ListElements(t *testing.T) {
	dll := algo.NewDoublyLinkedList()
	pokemon1 := domain.Pokemon{Id: 1, Name: "Pikachu"}
	pokemon2 := domain.Pokemon{Id: 2, Name: "Charmander"}

	dll.AddElementToEnd(1, pokemon1)
	dll.AddElementToEnd(2, pokemon2)

	pokemonList := dll.ListElements()

	if len(pokemonList) != 2 || pokemonList[0] != pokemon1 || pokemonList[1] != pokemon2 {
		t.Error("ListElements did not work as expected.")
	}
}
