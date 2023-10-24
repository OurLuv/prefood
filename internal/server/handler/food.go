package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/OurLuv/prefood/internal/model"
)

// * Get all
func (h *Handler) GetAllFood(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/food.html", "static/header.html")
	if err != nil {
		errStr := fmt.Sprintf("cannot find a template: %s", err.Error())
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}
	mp := make(map[string]interface{})
	var food []model.Food
	if food, err = h.service.FoodService.GetAll(); err != nil {
		errStr := fmt.Sprintf("storage error: %s", err.Error())
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}
	mp["Food"] = food
	mp["Title"] = "Your menu"
	err = tmpl.Execute(w, mp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// * Get food by id
func (h *Handler) GetFoodById(w http.ResponseWriter, r *http.Request) {

}

// * Create food
func (h *Handler) CreateFood(w http.ResponseWriter, r *http.Request) {
	var food model.Food
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&food); err != nil {
		errStr := fmt.Sprintf("can't validate: %s", err.Error())
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": errStr,
		})
		return
	}
	if err := h.service.FoodService.Create(food); err != nil {
		errStr := fmt.Sprintf("can't validate: %s", err.Error())
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "can't create row in database: " + errStr,
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Row is created!",
	})
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
