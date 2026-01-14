package models

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
	Stations map[string]Station `json:"stations"`
}

type Station struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Distance  float64 `json:"distance"`
}
