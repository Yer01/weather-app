package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// func (app *application) getAll(w http.ResponseWriter, r *http.Request) {

// }

func (app *application) getToday(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")
	country := chi.URLParam(r, "country")
	days := r.PathValue("days")

	if city == "" || country == "" {
		http.Error(w, "Missing fields", http.StatusBadRequest)
		app.logger.Println("Requesting with missing fields")
		return
	}

	if days == "" {
		days = "7"
	}

	res, err := app.weatherdata.Get(country, city, days)

	if err != nil {
		app.logger.Println(err.Error())
		http.Error(w, "Problems with data fetching", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(res); err != nil {
		app.logger.Println(err.Error())
		http.Error(w, "Problems with data encoding", http.StatusInternalServerError)
		return
	}
}
