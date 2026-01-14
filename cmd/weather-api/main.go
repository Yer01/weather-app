package main

import (
	"log"
	"os"

	"github.com/Yer01/weather-app/internal/api/routes"
	"github.com/Yer01/weather-app/internal/config"
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

	if err = routes.Routes(cfg); err != nil {
		log.Fatalf("Can't launch server on port 8081: %v", err)
	}
	os.Exit(0)
}
