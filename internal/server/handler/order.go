package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/gorilla/mux"
)

// * Create order
func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	order := model.Order{}
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Can't decode order: "+err.Error(), http.StatusInternalServerError)
		return
	}
	RId := mux.Vars(r)["id"]
	u64, _ := strconv.ParseUint(RId, 10, 32)
	restaurantId := uint(u64)
	order.RestaurantId = restaurantId
	if err := h.service.OrderService.Create(order); err != nil {
		http.Error(w, "Can't create order: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// * Get all orders
func (h *Handler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	RId := mux.Vars(r)["restaurant_id"]
	u64, _ := strconv.ParseUint(RId, 10, 32)
	restaurantId := uint(u64)

	orders, err := h.service.OrderService.GetAll(restaurantId)
	if err != nil {
		http.Error(w, "can't get order: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&orders)
}

// * Get by id
func (h *Handler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	RId := mux.Vars(r)["restaurant_id"]
	u64, _ := strconv.ParseUint(RId, 10, 32)
	restaurantId := uint(u64)
	OId := mux.Vars(r)["order_id"]
	u64, _ = strconv.ParseUint(OId, 10, 32)
	orderId := uint(u64)

	order, err := h.service.OrderService.GetById(restaurantId, orderId)
	if err != nil {
		http.Error(w, "can't get order: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(&order)
}
