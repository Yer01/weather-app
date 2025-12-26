package models

import (
	"fmt"
	"io"
	"net/http"
)

type WeatherData struct {
	Address  string `json:"address"`
	TimeZone string `json:"timezone"`
	APIKey   string
}

func (wd *WeatherData) get(country string, city string) (string, error) {
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s,%s?key=%s", city, country, wd.APIKey)

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

	return string(body), nil
}
