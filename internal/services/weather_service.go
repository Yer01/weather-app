package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Yer01/weather-app/internal/cache"
	"github.com/Yer01/weather-app/internal/models"
)

type Service interface {
	GetWeather(string, string) (models.WeatherData, error)
}

type weatherService struct {
	APIkey string
	cache  cache.Cache
}

func NewService(APIkey string, cache cache.Cache) Service {
	return &weatherService{
		APIkey: APIkey,
		cache:  cache,
	}
}

func (ws *weatherService) GetWeather(city string, country string) (models.WeatherData, error) {
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s,%s?key=%s", city, country, ws.APIkey)

	if city == "" || country == "" {
		log.Print("Trying to make request with missing fields")
		return models.WeatherData{}, fmt.Errorf("Trying to make request with missing fields")
	}

	cacheKey := city + country

	cachedData, err := ws.cache.Get(context.TODO(), cacheKey)

	if cachedData.Address != "" && cachedData.TimeZone != "" {
		return cachedData, nil
	}

	if err != nil {
		log.Printf("Error during cached data fetch: %v", err)
		return models.WeatherData{}, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error in creating new request: %v", err)
		return models.WeatherData{}, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Printf("Problem with making request: %v", err)
		return models.WeatherData{}, err
	}

	defer res.Body.Close()

	var data models.WeatherData

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Printf("Problem with reading request body: %v", err)
		return models.WeatherData{}, err
	}

	if err = json.Unmarshal(body, &data); err != nil {
		log.Printf("Problem with decoding data from response body: %v", err)
		return models.WeatherData{}, err
	}

	if err = ws.cache.Set(context.TODO(), cacheKey, data, 12*time.Hour); err != nil {
		return models.WeatherData{}, err
	}

	return data, nil
}
