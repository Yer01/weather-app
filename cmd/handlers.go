package main

import (
	"net/http"
)

// func (app *application) getAll(w http.ResponseWriter, r *http.Request) {

// }

func (app *application) getSingle(w http.ResponseWriter, r *http.Request) {
	city := r.FormValue("city")
	country := r.FormValue("country")
}
