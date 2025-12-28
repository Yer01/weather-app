package models

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type WeatherData struct {
	Address     string `json:"address"`
	TimeZone    string `json:"timezone"`
	APIKey      string
	RedisClient *redis.Client
}

func (wd *WeatherData) Get(country string, city string) (string, error) {
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s,%s?key=%s", city, country, wd.APIKey)

	key := country + city
	resp := ""

	cont, err := wd.RedisClient.Get(context.TODO(), key).Result()
	if err == redis.Nil {
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return "", err
		}

		req.Header.Set("Accept", "application/json")

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			return "", err
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)

		if err != nil {
			return "", err
		}
		resp = string(body)
		wd.RedisClient.Set(context.TODO(), key, resp, time.Hour*12)
	} else if err != nil {
		return "", err
	} else {
		resp = cont
	}
	return resp, nil
}
