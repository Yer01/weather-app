package cache

import "github.com/redis/go-redis/v9"

type Cache interface {
	Set()
	Get()
}

type redisCache struct {
	cache *redis.Client
}

func NewCache() Cache {
	rc := redis.NewClient(&redis.Options{})

	return &redisCache{
		cache: rc,
	}
}

func (r *redisCache) Set() {

}

func (r *redisCache) Get() {

}
