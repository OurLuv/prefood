package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetAllRestaurants(w http.ResponseWriter, r *http.Request) {
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
func (h *Handler) CreateRestaurant(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) GetRestaurantById(w http.ResponseWriter, r *http.Request) {

}
