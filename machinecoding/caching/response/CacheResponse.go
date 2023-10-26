package response

import "caching/domain"

type CacheResponse struct {
	Pokemon domain.Pokemon
	Err     error
}

func Init(Err error) CacheResponse {
	return CacheResponse{domain.Pokemon{}, Err}
}
