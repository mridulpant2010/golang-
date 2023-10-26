package cache

import (
	algo "caching/algorithms"
	domain "caching/domain"
	response "caching/response"
	"fmt"
)

type LRUCaching struct {
	Capacity int
	Storage  map[int]*algo.Node //what map shall store as value?
	Dll      algo.DoublyLinkedList
}

func InItCache(capacity int) *LRUCaching {

	return &LRUCaching{
		Capacity: capacity,
		Storage:  make(map[int]*algo.Node, capacity),
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
	//if the key already exists update its value otherwise append to end
	if self.Dll.Len >= self.Capacity {
		//perform eviction
		self.Evict()
	}
	node, exist := self.Storage[key]
	if exist {
		//if the key already exists removing that node and adding it to the last.
		self.Dll.DetachNode(node)
	}
	self.Storage[key] = self.Dll.AddElementToEnd(key, value)

}

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
	delete(self.Storage,key)
	return response.CacheResponse{Pokemon: node.Data.Value}
}
