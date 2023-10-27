package response

import entity "pokemon.com/cache/src/entity"

type CacheResponse struct {
	Pokemon entity.Pokemon
	Err     error
}

func Init(Err error) CacheResponse {
	return CacheResponse{entity.Pokemon{}, Err}
}
