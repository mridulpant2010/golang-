package main

import (
	ca "caching/cache"
)

func main() {
	cacheObj := ca.InItCache(5)
	//fmt.Println(cacheObj.Capacity, cacheObj.Dll, cacheObj.Storage)

	cacheObj.Put(1, "one")
	cacheObj.Put(2, "two")
	cacheObj.Put(3, "three")
	cacheObj.Put(4, "four")
	cacheObj.Put(5, "five")
	cacheObj.Put(6, "six")
	cacheObj.ListAll()

	cacheObj.Dll.Display()
	cacheObj.Get(2)
	cacheObj.Dll.Display()
	cacheObj.Put(3, "ndsda")
	cacheObj.Dll.Display()
	cacheObj.Get(1)
}
