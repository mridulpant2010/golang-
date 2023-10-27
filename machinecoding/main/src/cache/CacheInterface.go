package cache

import (
	entity "pokemon.com/cache/src/entity"
	response "pokemon.com/cache/src/response"
)

type Caching interface {
	Get(key int) response.CacheResponse
	Put(key int, value entity.Pokemon)
	Evict()
	ListAll() []entity.Pokemon
	DeleteById(key int) response.CacheResponse
	Length() int
}
