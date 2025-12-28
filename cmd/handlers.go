package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// func (app *application) getAll(w http.ResponseWriter, r *http.Request) {

// }

func (app *application) getSingle(w http.ResponseWriter, r *http.Request) {
	city := r.FormValue("city")
	country := r.FormValue("country")

	res, err := app.weatherdata.Get(country, city)

	if err != nil {
		fmt.Println("HEREREREREREERERE")
		app.logger.Println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(res); err != nil {
		app.logger.Println(err.Error())
		return
	}
}
