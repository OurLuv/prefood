package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/gorilla/mux"
)

// * Create
func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Сategory
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CategoryService.Create(category)
	if err != nil {
		http.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusCreated)
}

// * Read by id
func (h *Handler) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categoryID := mux.Vars(r)["id"]
	u64, _ := strconv.ParseUint(categoryID, 10, 32)
	id := uint(u64)
	category, err := h.service.CategoryService.GetById(id)
	if err != nil {
		http.Error(w, "Failed to get category: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

// * Read all
func (h *Handler) GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	category, err := h.service.CategoryService.GetAll()
	if err != nil {
		http.Error(w, "Failed to get category", http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(category)
}

// * Update category by id
func (h *Handler) UpdateCategoryById(w http.ResponseWriter, r *http.Request) {
	var category model.Сategory
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.service.CategoryService.Create(category)
	if err != nil {
		http.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusCreated)
}

// * Delete by id
func (h *Handler) DeleteCategoryById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categoryID := mux.Vars(r)["id"]
	u64, _ := strconv.ParseUint(categoryID, 10, 32)
	id := uint(u64)
	err := h.service.CategoryService.DeleteCategoryById(id)
	if err != nil {
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
