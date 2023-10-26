package algorithms

import (
	"fmt"
)

type Pair struct {
	Key   int
	Value string
}

func CreatePair(key int, value string) *Pair {
	return &Pair{key, value}
}

type Node struct {
	Prev *Node
	Next *Node
	Data *Pair
}

func InItNode(key int, value string) *Node {
	return &Node{nil, nil, CreatePair(key, value)}
}

type DoublyLinkedList struct {
	Tail *Node
	Head *Node
	Len  int
}

// create an empty constructor
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil, nil, 0}
}

func (self *DoublyLinkedList) AddElementToEnd(key int, value string) *Node {
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

func (dll *DoublyLinkedList) RemoveFromBeginning() *Node {
	if dll.Head == nil {
		fmt.Println("Element doesn't exist")
		return nil
	}

	temp := dll.Head
	dll.Head = dll.Head.Next
	if dll.Head != nil {
		dll.Head.Prev = nil
	} else {
		dll.Tail = nil
	}

	dll.Len--
	return temp
}

func (dll *DoublyLinkedList) MoveToEnd(node *Node) {
	if node != nil {
		// Detach the node from its current position
		if node.Prev != nil {
			node.Prev.Next = node.Next
		} else {
			dll.Head = node.Next
		}
		if node.Next != nil {
			node.Next.Prev = node.Prev
		} else {
			dll.Tail = node.Prev
		}

		// Add the node to the end of the list
		node.Prev = dll.Tail
		node.Next = nil
		if dll.Tail != nil {
			dll.Tail.Next = node
		} else {
			dll.Head = node
		}
		dll.Tail = node
	}
}

func (dll *DoublyLinkedList) DetachNode(node *Node) {
	if node != nil {
		if node.Prev != nil {
			node.Prev.Next = node.Next
		} else {
			dll.Head = node.Next
		}
		if node.Next != nil {
			node.Next.Prev = node.Prev
		} else {
			dll.Tail = node.Prev
		}
		dll.Len--
	}
}
func (self *DoublyLinkedList) Display() {

	temp := self.Head
	for temp != nil {
		fmt.Println(temp.Data.Key, temp.Data.Value)
		//fmt.Println(temp)
		temp = temp.Next
	}
}

func (self *DoublyLinkedList) Size() int {
	return self.Len
}
