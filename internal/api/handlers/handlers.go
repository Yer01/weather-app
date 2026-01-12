package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Yer01/weather-app/internal/services"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service services.Service
}

func NewHandler(service services.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetToday(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")
	country := chi.URLParam(r, "country")
	// days := r.PathValue("days")

	data, err := h.service.GetWeather(city, country)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(data); err != nil {
		log.Println(err.Error())
		http.Error(w, "Problems with data encoding", http.StatusInternalServerError)
		return
	}
}
