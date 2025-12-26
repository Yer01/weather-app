package main

import (
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
	mux := chi.NewRouter()

	//mux.Get("/", app.getAll)
	mux.Get("/report", app.getSingle)

	return mux
}
