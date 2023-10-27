package algorithms

import (
	"fmt"

	entity "pokemon.com/cache/src/entity"
)

type Pair struct {
	Key   int
	Value entity.Pokemon
}

// initialize a constructor for pair
func CreatePair(key int, value entity.Pokemon) *Pair {
	return &Pair{key, value}
}

// create a Node which stores prev and next pointer and key-value pair data
type Node struct {
	Prev *Node
	Next *Node
	Data *Pair
}

// initialize the constructor for Node creation
func InItNode(key int, value entity.Pokemon) *Node {
	return &Node{nil, nil, CreatePair(key, value)}
}

// head and tail are pointers for starting and the ending position
type DoublyLinkedList struct {
	Tail *Node
	Head *Node
	Len  int
}

// INFO: initialize empty constructor for the Dll.
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil, nil, 0}
}

// INFO: adds a new node to the end of the DLL.
func (self *DoublyLinkedList) AddElementsToEnd(key int, value entity.Pokemon) *Node {
	newNode := InItNode(key, value)
	if self.Head == nil {
		self.Head = newNode
		self.Tail = newNode
	} else {
		self.Tail.Next = newNode
		newNode.Prev = self.Tail
		self.Tail = self.Tail.Next
	}
	self.Len += 1
	return newNode
}

// removes a node from the beginning of the DLLs
func (self *DoublyLinkedList) RemoveFromBeginning() *Node {
	if self.Head == nil {
		fmt.Println("Element doesn't exists")
		return nil
	}
	temp := self.Head
	self.Head = self.Head.Next
	if self.Head != nil {
		self.Head.Prev = nil
	} else {
		self.Tail = nil
	}
	self.Len -= 1
	return temp
}

// detach the node from its current position and move to the end of the DLL.
func (self *DoublyLinkedList) MoveToEnd(node *Node) {
	if node != nil {
		self.DetachNode(node)

		//add node to the end of the DLL
		node.Prev = self.Tail
		node.Next = nil
		if self.Tail != nil {
			self.Tail.Next = node
		} else {
			self.Head = node
		}
		self.Tail = node
	}

}

// detach a given node from its current position.
func (self *DoublyLinkedList) DetachNode(node *Node) {
	if node != nil {
		if node.Prev != nil {
			node.Prev.Next = node.Next
		} else {
			self.Head = node.Next
		}
		if node.Next != nil {
			node.Next.Prev = node.Prev
		} else {
			self.Tail = node.Prev
		}
		self.Len--
	}
}

// display all the linkedlist nodes
func (self *DoublyLinkedList) Display() {
	temp := self.Head
	for temp != nil {
		fmt.Println(temp.Data.Key, temp.Data.Value)
		temp = temp.Next
	}

}

// prepares a list of pokemon items according to their sequencing position
func (self *DoublyLinkedList) ListElements() []entity.Pokemon {
	var items []entity.Pokemon
	temp := self.Head
	for temp != nil {
		items = append(items, temp.Data.Value)
		temp = temp.Next
	}
	return items
}

func (self *DoublyLinkedList) Size() int {
	return self.Len
}
