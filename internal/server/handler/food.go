package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/OurLuv/prefood/internal/model"
)

// * Get all
func (h *Handler) GetAllFood(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/food.html")
	if err != nil {
		errStr := fmt.Sprintf("cannot find a template: %s", err.Error())
		http.Error(w, errStr, http.StatusBadRequest)
		return
	}
	mp := make(map[string][]model.Food)
	var food []model.Food
	if food, err = h.service.FoodService.GetAll(); err != nil {
		http.Error(w, "storage error", http.StatusBadRequest)
		return
	}
	mp["Food"] = food
	tmpl.Execute(w, mp)

}
