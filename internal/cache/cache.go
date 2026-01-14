package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Yer01/weather-app/internal/models"
	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Set(context.Context, string, models.WeatherData, time.Duration) error
	Get(context.Context, string) (models.WeatherData, error)
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

func (r *redisCache) Set(ctx context.Context, key string, data models.WeatherData, expirationTime time.Duration) error {
	cachedata, err := json.Marshal(data)

	if err != nil {
		log.Printf("Error with marshaling cache data: %v", err)
		return err
	}

	if err := r.cache.Set(ctx, key, cachedata, expirationTime).Err(); err != nil {
		log.Printf("Problem with setting data to cache: %v", err)
		return err
	}

	return nil
}

func (r *redisCache) Get(ctx context.Context, key string) (models.WeatherData, error) {
	data, err := r.cache.Get(ctx, key).Result()

	if err != nil {
		log.Print("Cache Miss!")
		return models.WeatherData{}, err
	}

	if data == "" {
		return models.WeatherData{}, err
	}

	var weatherdata models.WeatherData

	if err = json.Unmarshal([]byte(data), &weatherdata); err != nil {
		log.Printf("Error during decoding of cached data: %v", err)
		return models.WeatherData{}, err
	}
	log.Print("Cache Hit!!!")
	return weatherdata, nil
}
