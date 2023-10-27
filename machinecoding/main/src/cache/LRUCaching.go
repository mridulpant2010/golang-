package cache

import (
	"fmt"

	algo "pokemon.com/cache/src/algorithms"
	domain "pokemon.com/cache/src/entity"
	response "pokemon.com/cache/src/response"
)

type LRUCaching struct {
	Capacity int
	Storage  map[int]*algo.Node //key as int and value as a pointer to DLL node.
	Dll      algo.DoublyLinkedList
}

// initialize a new lru cache with the specified capacity
func InItCache(capacity int) Caching {

	return &LRUCaching{
		Capacity: capacity,
		Storage:  make(map[int]*algo.Node, capacity), //Storage is the hashMap
		Dll:      *algo.NewDoublyLinkedList(),
	}

}

func (self *LRUCaching) Get(key int) response.CacheResponse {
	node, exists := self.Storage[key]

	if !exists {
		//node exists so we will have to move it to last
		err := fmt.Errorf("Key doesnt exist")
		return response.Init(err)
	}
	//print(node.Data.Key, node.Data.Value)
	self.Dll.MoveToEnd(node)
	return response.CacheResponse{Pokemon: node.Data.Value}

}

func (self *LRUCaching) Put(key int, value domain.Pokemon) {

	//perform eviction first based on LRU eviction polcy
	if self.Dll.Len >= self.Capacity {
		//perform eviction
		self.Evict()
	}
	node, exist := self.Storage[key]
	if exist {
		//if the key already exists removing that node
		self.Dll.DetachNode(node)
	}
	// and adding it to the last.
	self.Storage[key] = self.Dll.AddElementsToEnd(key, value)

}

// INFO: eviction removes the least used node from the beginning of the DLL and delete it from the hashMap as well.
func (self *LRUCaching) Evict() {
	node := self.Dll.RemoveFromBeginning()
	delete(self.Storage, node.Data.Key)
}

func (self *LRUCaching) ListAll() []domain.Pokemon {

	return self.Dll.ListElements()
}

func (self *LRUCaching) DeleteById(key int) response.CacheResponse {
	node, exist := self.Storage[key]
	if !exist {
		//node exists so we will have to move it to last
		err := fmt.Errorf("Key doesnt exist")
		return response.Init(err)
	}
	self.Dll.DetachNode(node)
	delete(self.Storage, key)
	return response.CacheResponse{Pokemon: node.Data.Value}
}

func (self *LRUCaching) Length() int { return self.Dll.Len }
