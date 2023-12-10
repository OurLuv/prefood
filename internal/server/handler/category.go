package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// * Create
func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var category model.Сategory
	// getting data from request
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		h.logger.Error("Can't get id from context", err)
		SendError(w, "There is no data", 400)
		return
	}
	// setting restaurant id from context
	restaurant, ok := r.Context().Value("restaurant").(*model.Restaurant)
	if !ok {
		h.logger.Error("Can't get restaurant from context")
		SendError(w, "Can't get a restauarant", http.StatusInternalServerError)
		return
	}
	category.RestaurantId = restaurant.Id

	// validatation
	if err := validator.New().Struct(category); err != nil {
		h.logger.Error("validation err: ", err)
		resp := ValidateError(err.(validator.ValidationErrors))
		SendRespError(w, resp, 400)
		return
	}

	// creating category
	err = h.service.CategoryService.Create(category)
	if err != nil {
		h.logger.Error("Can't create category", err)
		SendError(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	// sending response
	resp := Response{
		Success: true, Message: "Category is added",
	}
	json.NewEncoder(w).Encode(resp)
}

// * Get by id
func (h *Handler) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// getting category id from url
	categoryID := mux.Vars(r)["category_id"]
	u64, _ := strconv.ParseUint(categoryID, 10, 32)
	id := uint(u64)

	// looking for category
	category, err := h.service.CategoryService.GetById(id)
	if err != nil {
		h.logger.Error("Can't get category from db", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

// * Get all
func (h *Handler) GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// getting restaurant id from url
	restaurantId := mux.Vars(r)["restaurant_id"]
	u64, _ := strconv.ParseUint(restaurantId, 10, 32)
	id := uint(u64)

	// looking for all categories
	category, err := h.service.CategoryService.GetAll(id)
	if err != nil {
		h.logger.Error("Can't get categories from db", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

// * Update category by id
func (h *Handler) UpdateCategoryById(w http.ResponseWriter, r *http.Request) {
	var category model.Сategory

	// getting data from request
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		h.logger.Error("no date from request", err)
		SendError(w, "There is no date", http.StatusBadRequest)
		return
	}

	// getting category id from url
	categoryId := mux.Vars(r)["category_id"]
	u64, _ := strconv.ParseUint(categoryId, 10, 32)
	id := uint(u64)
	category.Id = id

	// validation
	if err := validator.New().Struct(category); err != nil {
		h.logger.Error("validation err: ", err)
		resp := ValidateError(err.(validator.ValidationErrors))
		SendRespError(w, resp, 400)
		return
	}

	// updating category
	err = h.service.CategoryService.UpdateCategory(category)
	if err != nil {
		h.logger.Error("Can't get categories from db", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}

	// sending response
	resp := Response{
		Success: true, Message: "Category is updated",
	}
	json.NewEncoder(w).Encode(resp)
}

// * Delete by id
func (h *Handler) DeleteCategoryById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// getting categoryId & restaurantId
	categoryID := mux.Vars(r)["category_id"]
	u64, _ := strconv.ParseUint(categoryID, 10, 32)
	id := uint(u64)

	// deleting category
	err := h.service.CategoryService.DeleteCategoryById(id)
	if err != nil {
		h.logger.Error("Can't delete categories from db", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}

	// sending response
	resp := Response{
		Success: true, Message: "Category is deleted",
	}
	json.NewEncoder(w).Encode(resp)
}
