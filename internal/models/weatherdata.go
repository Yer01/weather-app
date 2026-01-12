package models

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type WeatherData struct {
	Address  string  `json:"address"`
	TimeZone string  `json:"timezone"`
	TZOffset float64 `json:"tzoffset"`
	Days     []struct {
		Datetime   string  `json:"datetime"` // YYYY-MM-DD
		Temp       float64 `json:"temp"`
		Humidity   float64 `json:"humidity"`
		WindSpeed  float64 `json:"windspeed"`
		Pressure   float64 `json:"pressure"`
		CloudCover float64 `json:"cloudcover"`
		Sunrise    string  `json:"sunrise"`
		Sunset     string  `json:"sunset"`
	} `json:"days"`
	Strations   map[string]Station `json:"stations"`
	APIKey      string             `json:"-"`
	RedisClient *redis.Client      `json:"-"`
}

type Station struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Distance  float64 `json:"distance"`
}

func (wd *WeatherData) Get(country string, city string, days string) (WeatherData, error) {
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s,%s?key=%s", city, country, wd.APIKey)

	key := country + city + days
	var data WeatherData
	daysNum, _ := strconv.Atoi(days)
	cont, err := wd.RedisClient.Get(context.TODO(), key).Bytes()
	if err == redis.Nil {
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			log.Println("Problems with request sent")
			return WeatherData{}, err
		}

		req.Header.Set("Accept", "application/json")

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			log.Println("Problems with making request")
			return WeatherData{}, err
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)

		if err != nil {
			log.Println("Problems with read of the data")
			return WeatherData{}, err
		}

		if err = json.Unmarshal(body, &data); err != nil {
			log.Println("Problems with unmarshaling data")
			return WeatherData{}, err
		}

		if daysNum > 0 && daysNum < len(data.Days) {
			data.Days = data.Days[:daysNum]
		} else if daysNum <= 0 || daysNum >= len(data.Days) {
			log.Print("Invalid days value")
			return WeatherData{}, fmt.Errorf("Invalid days value")
		}

		wd.RedisClient.Set(context.TODO(), key, body, time.Hour*12)
	} else if err != nil {
		log.Println("Other problem with Redis server")
		return WeatherData{}, err
	} else {
		log.Print("Fetched data from redis")
		if err = json.Unmarshal(cont, &data); err != nil {
			log.Println("Problems with unmarshaling data")
			return WeatherData{}, err
		}
		if daysNum > 0 && daysNum < len(data.Days) {
			data.Days = data.Days[:daysNum]
		} else if daysNum <= 0 || daysNum >= len(data.Days) {
			log.Print("Invalid days value")
			return WeatherData{}, fmt.Errorf("Invalid days value")
		}
	}
	return data, nil
}
