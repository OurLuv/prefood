package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type FoodResponse struct {
	Response Response     `json:"response"`
	Menu     []model.Food `json:"menu,omitempty"`
	Food     *model.Food  `json:"food,omitempty"`
}

// * Get all
func (h *Handler) GetAllFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//getting data from url
	RId := mux.Vars(r)["restaurant_id"]
	u64, _ := strconv.ParseUint(RId, 10, 32)
	restaurantId := uint(u64)

	// looking for all restaurants
	var food []model.Food
	var err error
	if food, err = h.service.FoodService.GetAll(restaurantId); err != nil {
		h.logger.Error("storage error: ", err)
		SendError(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	// sending response
	resp := FoodResponse{
		Response: Response{Success: true},
		Menu:     food,
	}
	json.NewEncoder(w).Encode(resp)

}

// * Get food by id
func (h *Handler) GetFoodById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//getting data from url
	RId := mux.Vars(r)["restaurant_id"]
	u64, _ := strconv.ParseUint(RId, 10, 32)
	restaurantId := uint(u64)
	FId := mux.Vars(r)["id"]
	u64, _ = strconv.ParseUint(FId, 10, 32)
	foodId := uint(u64)

	// looking for all restaurants
	var food *model.Food
	var err error
	if food, err = h.service.FoodService.GetById(restaurantId, foodId); err != nil {
		h.logger.Error("storage error: ", err)
		SendError(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	// sending response
	resp := FoodResponse{
		Response: Response{Success: true},
		Food:     food,
	}
	json.NewEncoder(w).Encode(resp)
}

// * Create food
func (h *Handler) CreateFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var food model.Food
	// getting data from request
	if err := json.NewDecoder(r.Body).Decode(&food); err != nil {
		h.logger.Error("can't get data from request", err)
		SendError(w, "There is no data", http.StatusBadRequest)
		return
	}

	// setting restaurant id from context
	restaurant, ok := r.Context().Value("restaurant").(*model.Restaurant)
	if !ok {
		h.logger.Error("Can't get restaurant from context")
		SendError(w, "Can't get a restauarant", http.StatusInternalServerError)
		return
	}
	food.RestaurantId = restaurant.Id

	// validatation
	if err := validator.New().Struct(food); err != nil {
		h.logger.Error("validation err: ", err)
		resp := ValidateError(err.(validator.ValidationErrors))
		SendRespError(w, resp, 400)
		return
	}

	// creating food
	if err := h.service.FoodService.Create(food); err != nil {
		h.logger.Error("can't create food", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}
	resp := Response{Success: true, Message: "Row is added to database"}
	json.NewEncoder(w).Encode(resp)
}

// * Update
func (h *Handler) UpdateFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var food model.Food
	// getting data from request
	if err := json.NewDecoder(r.Body).Decode(&food); err != nil {
		h.logger.Error("can't get data from request", err)
		SendError(w, "There is no data", http.StatusBadRequest)
		return
	}
	FId := mux.Vars(r)["id"]
	u64, _ := strconv.ParseUint(FId, 10, 32)
	id := uint(u64)
	food.Id = id

	// validatation
	if err := validator.New().Struct(food); err != nil {
		h.logger.Error("validation err: ", err)
		resp := ValidateError(err.(validator.ValidationErrors))
		SendRespError(w, resp, 400)
		return
	}

	// creating food
	if err := h.service.FoodService.UpdateById(food); err != nil {
		h.logger.Error("can't create food", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}
	resp := Response{Success: true, Message: "Item is updated"}
	json.NewEncoder(w).Encode(resp)
}

// * Delete
func (h *Handler) DeleteFood(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) CreateFoodView(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/create-food.html")
	if err != nil {
		errStr := fmt.Sprintf("can't load a view 01: %s", err.Error())
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}
	mp := make(map[string][]model.Сategory)
	if mp["Categories"], err = h.service.CategoryService.GetAll(); err != nil {
		fmt.Print(err.Error())
	}

	if err = t.Execute(w, mp); err != nil {
		errStr := fmt.Sprintf("can't load a view 02: %s", err.Error())
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}
}
