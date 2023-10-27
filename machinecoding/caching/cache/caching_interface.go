package cache

import (
	domain "caching/domain"
	"caching/response"
)

type Caching interface {
	Get(key int) response.CacheResponse
	Put(key int, value domain.Pokemon)
	Evict()
	ListAll()
	DeleteById(key int) response.CacheResponse
}
