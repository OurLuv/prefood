package handler

import (
	"encoding/json"
	"net/http"

	"github.com/OurLuv/prefood/internal/model"
)

// * Get all
func (h *Handler) GetAllRestaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, ok := r.Context().Value("id").(uint)
	if !ok {
		http.Error(w, "Can't get restaurants 1: ", http.StatusInternalServerError)
		return
	}
	models, err := h.service.RestaruantService.GetAll(id)
	if err != nil {
		http.Error(w, "Can't get restaurants 2: "+err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(models)
}

// * Create
func (h *Handler) CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("id").(uint)
	if !ok {
		http.Error(w, "Can't get your id", http.StatusInternalServerError)
		return
	}
	var restaurant model.Restaurant
	if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
		http.Error(w, "Can't get data: "+err.Error(), http.StatusInternalServerError)
		return
	}
	restaurant.ClientId = id
	err := h.service.RestaruantService.Create(restaurant)
	if err != nil {
		http.Error(w, "Can't create a restaurant: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// * Get by id
func (h *Handler) GetRestaurantById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurant, ok := r.Context().Value("restaurant").(*model.Restaurant)
	if !ok {
		http.Error(w, "Can't get restaurant from context", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&restaurant)
}
