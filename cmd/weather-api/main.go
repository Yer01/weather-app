package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Yer01/weather-app/internal/api/handlers"
	"github.com/Yer01/weather-app/internal/api/routes"
	"github.com/Yer01/weather-app/internal/cache"
	"github.com/Yer01/weather-app/internal/config"
	"github.com/Yer01/weather-app/internal/services"
	"github.com/joho/godotenv"
)

func main() {

	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatalf("Couldnt load environmental variables: %v", err)
	}

	apikey := os.Getenv("API_KEY")

	cfg := config.Config{
		APIkey:     apikey,
		ServerPort: "8081",
	}

	log.Print("Starting server on port 8081...")

	weatherCache := cache.NewCache()

	weatherService := services.NewService(cfg.APIkey, weatherCache)

	weatherHandler := handlers.NewHandler(weatherService)

	router := routes.Routes(*weatherHandler)

	if err = http.ListenAndServe(fmt.Sprintf("localhost:%s", cfg.ServerPort), router); err != nil {
		log.Fatalf("Can't launch server on port 8081: %v", err)
	}
	os.Exit(0)
}
