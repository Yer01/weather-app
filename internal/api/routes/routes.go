package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Yer01/weather-app/internal/api/handlers"
	"github.com/Yer01/weather-app/internal/cache"
	"github.com/Yer01/weather-app/internal/config"
	"github.com/Yer01/weather-app/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func Routes(cfg config.Config) error {
	mux := chi.NewRouter()

	weatherCache := cache.NewCache()

	weatherService := services.NewService(cfg.APIkey, weatherCache)

	weatherHandler := handlers.NewHandler(weatherService)

	mux.Use(middleware.Logger)

	mux.Use(httprate.Limit(100, 10*time.Second, httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint)))

	mux.Get("/report/{country}/{city}", weatherHandler.GetToday)

	err := http.ListenAndServe(fmt.Sprintf("localhost:%s", cfg.ServerPort), mux)

	return err
}
