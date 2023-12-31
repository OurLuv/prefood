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
// @Summary GetRestaurants
// @Security ApiKeyAuth
// @Tags Restaurant
// @Description get all restaurants
// @ID get-restaurants
// @Produce json
// @Failure default {object} Response
// @Router /restaurants [get]
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
		Response:    Response{Status: 1},
		Restaurants: models,
	}
	json.NewEncoder(w).Encode(resp)
}

// * Create
// @Summary CreateRestaurant
// @Security ApiKeyAuth
// @Tags Restaurant
// @Description create restaurant
// @ID create-restaurant
// @Accept json
// @Produce json
// @Param input body model.Restaurant true "restaurant info"
// @Success 200 {object} ResponseId
// @Failure default {object} Response
// @Router /restaurants [post]
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
	restaurant_id, err := h.service.RestaruantService.Create(restaurant)
	if err != nil {
		h.logger.Error("Can't get data:", err)
		SendError(w, "There is no data", http.StatusInternalServerError)
		return
	}
	resp := ResponseId{
		Response: Response{
			Status: 1,
		},
		Id: restaurant_id,
	}
	json.NewEncoder(w).Encode(&resp)
}

// * Update
// @Summary UpdateRestaurant
// @Security ApiKeyAuth
// @Tags Restaurant
// @Description update restaurant
// @ID update-restaurant
// @Param restaurant_id path int true "restaurant id"
// @Accept json
// @Produce json
// @Param input body model.Restaurant true "restaurant info"
// @Success 200 {object} Response
// @Failure default {object} Response
// @Router /restaurants/{restaurant_id} [put]
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
		Status: 1,
	}
	json.NewEncoder(w).Encode(&resp)
}

// * Get by id
// @Summary GetRestaurantById
// @Security ApiKeyAuth
// @Tags Restaurant
// @Description get restaurant by id
// @ID get-restaurant-by-id
// @Param restaurant_id path int true "restaurant id"
// @Produce json
// @Success 200 {object} ResponseRestaurant
// @Failure default {object} Response
// @Router /restaurants/{restaurant_id} [get]
func (h *Handler) GetRestaurantById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurant, ok := r.Context().Value("restaurant").(*model.Restaurant)
	if !ok {
		h.logger.Error("Can't get restaurant from context", http.StatusInternalServerError)
		SendError(w, "Can't get a restauarant", http.StatusInternalServerError)
		return
	}

	_ = restaurant
	resp := ResponseRestaurant{
		Response:   Response{Status: 1},
		Restaurant: restaurant,
	}
	json.NewEncoder(w).Encode(&resp)
}

// * Delete
// @Summary DeleteRestaurant
// @Security ApiKeyAuth
// @Tags Restaurant
// @Description delete restaurant by id
// @ID delete-restaurant
// @Param restaurant_id path int true "restaurant id"
// @Produce json
// @Success 200 {object} Response
// @Failure default {object} Response
// @Router /restaurants/{restaurant_id} [delete]
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
		Status: 1,
	}
	json.NewEncoder(w).Encode(&resp)
}

// * Open-close
// @Summary OpenCloseRestaurant
// @Security ApiKeyAuth
// @Tags Restaurant
// @Description open or close restaurant
// @ID open-close-restaurant
// @Param restaurant_id path int true "restaurant id"
// @Produce json
// @Success 200 {object} Response
// @Failure default {object} Response
// @Router /restaurants/{restaurant_id}/openclose [post]
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

	_ = msg
	resp := Response{
		Status:  1,
		Message: msg,
	}
	json.NewEncoder(w).Encode(&resp)
}
