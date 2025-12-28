package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yer01/weather-app/internal/models"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type application struct {
	redisClient *redis.Client
	weatherdata *models.WeatherData
	logger      *log.Logger
}

func main() {

	logger := log.Default()

	rClient := redis.NewClient(&redis.Options{})

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apikey := os.Getenv("API_KEY")

	wd := &models.WeatherData{
		APIKey:      apikey,
		RedisClient: rClient,
	}

	app := &application{
		redisClient: rClient,
		weatherdata: wd,
		logger:      logger,
	}

	router := app.routes()
	logger.Print("starting server over localhost:8081...")
	if err := http.ListenAndServe("localhost:8081", router); err != nil {
		logger.Fatal(err.Error())
	}

	os.Exit(0)

}
