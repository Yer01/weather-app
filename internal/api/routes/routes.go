package routes

import (
	"time"

	"github.com/Yer01/weather-app/internal/api/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func Routes(weatherHandler *handlers.Handler) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)

	mux.Use(httprate.Limit(100, 10*time.Second, httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint)))

	mux.Get("/report/{country}/{city}", weatherHandler.GetToday)

	return mux
}
