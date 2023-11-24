package handler

import (
	"encoding/json"
	"net/http"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/go-playground/validator/v10"
)

type ResponseRestaurant struct {
	Response    Response
	Restaurant  *model.Restaurant  `json:"restaurant,omitempty"`
	Restaurants []model.Restaurant `json:"restaurants,omitempty"`
}

// * Get all
func (h *Handler) GetAllRestaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, ok := r.Context().Value("id").(uint)
	if !ok {
		h.logger.Error("can't get id from url")
		SendError(w, "Bad request", 400)
		return
	}
	models, err := h.service.RestaruantService.GetAll(id)
	if err != nil {
		h.logger.Error("can't get restaurants: ", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}
	resp := ResponseRestaurant{
		Response:    Response{Success: true},
		Restaurants: models,
	}
	json.NewEncoder(w).Encode(resp)
}

// * Create
func (h *Handler) CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// getting id from context
	id, ok := r.Context().Value("id").(uint)
	if !ok {
		h.logger.Error("Can't get id from context")
		SendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// getting data from request
	var restaurant model.Restaurant
	if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
		h.logger.Error("Can't get data:", err)
		SendError(w, "There is no data", http.StatusBadRequest)
		return
	}

	// validation
	if err := validator.New().Struct(restaurant); err != nil {
		h.logger.Error("validation err: ", err)
		resp := ValidateError(err.(validator.ValidationErrors))
		SendRespError(w, resp, 400)
		return
	}

	// creating restaurant
	restaurant.ClientId = id
	err := h.service.RestaruantService.Create(restaurant)
	if err != nil {
		h.logger.Error("Can't get data:", err)
		SendError(w, "There is no data", http.StatusInternalServerError)
		return
	}
	resp := Response{
		Success: true,
	}
	json.NewEncoder(w).Encode(&resp)
}

// * Update
func (h *Handler) UpdateRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// getting restaurant from context
	restaurant, ok := r.Context().Value("restaurant").(*model.Restaurant)
	if !ok {
		h.logger.Error("Can't get restaurant from context")
		SendError(w, "Can't get a restauarant", http.StatusInternalServerError)
		return
	}
	data := struct {
		Name    string `json:"name" validate:"required"`
		Phone   string `json:"phone" validate:"required"`
		Country string `json:"country"`
		State   string `json:"state"`
		City    string `json:"city"`
		Street  string `json:"street"`
	}{}
	// getting data from request
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		h.logger.Error("Can't get data:", err)
		SendError(w, "There is no data", http.StatusBadRequest)
		return
	}

	// validation
	if err := validator.New().Struct(data); err != nil {
		h.logger.Error("validation err: ", err)
		resp := ValidateError(err.(validator.ValidationErrors))
		SendRespError(w, resp, 400)
		return
	}

	// updating restaurant
	err := h.service.RestaruantService.Update(*restaurant, data)
	if err != nil {
		h.logger.Error("Can't get data:", err)
		SendError(w, "There is no data", http.StatusInternalServerError)
		return
	}

	resp := Response{
		Success: true,
	}
	json.NewEncoder(w).Encode(&resp)
}

// * Get by id
func (h *Handler) GetRestaurantById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurant, ok := r.Context().Value("restaurant").(*model.Restaurant)
	if !ok {
		h.logger.Error("Can't get restaurant from context", http.StatusInternalServerError)
		SendError(w, "Can't get a restauarant", http.StatusInternalServerError)
		return
	}
	resp := ResponseRestaurant{
		Response:   Response{Success: true},
		Restaurant: restaurant,
	}
	json.NewEncoder(w).Encode(&resp)
}

// * Delete
func (h *Handler) DeleteRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurant, ok := r.Context().Value("restaurant").(*model.Restaurant)
	if !ok {
		h.logger.Error("Can't get restaurant from context")
		SendError(w, "Can't get a restauarant", http.StatusInternalServerError)
		return
	}
	if err := h.service.RestaruantService.Delete(restaurant.Id); err != nil {
		h.logger.Error("Can't delete restaurant", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}
	resp := Response{
		Success: true,
	}
	json.NewEncoder(w).Encode(&resp)
}

func (h *Handler) OpenClose(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurant, ok := r.Context().Value("restaurant").(*model.Restaurant)
	if !ok {
		h.logger.Error("Can't get restaurant from context")
		SendError(w, "Can't get a restauarant", http.StatusInternalServerError)
		return
	}
	var open *bool
	var err error
	var msg string
	if open, err = h.service.RestaruantService.OpenClose(restaurant.Id); err != nil {
		h.logger.Error("Can't get restaurant", err)
		SendError(w, "Can't get a restauarant", http.StatusInternalServerError)
		return
	}
	if !*open {
		msg = "Restaurant is closed"
	} else {
		msg = "Restaurant is open"
	}
	resp := Response{
		Success: true,
		Message: msg,
	}
	json.NewEncoder(w).Encode(&resp)
}
