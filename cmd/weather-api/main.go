package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yer01/weather-app/internal/api/routes"
	"github.com/joho/godotenv"
)

func main() {

	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatalf("Couldnt load environmental variables: %v", err)
	}

	apikey := os.Getenv("API_KEY")

	router := routes.Routes(apikey)

	log.Print("Starting server on port 8081...")

	if err = http.ListenAndServe("localhost:8081", router); err != nil {
		log.Fatalf("Can't launch server on port 8081: %v", err)
	}
	os.Exit(0)
}
