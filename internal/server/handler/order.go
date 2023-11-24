package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type ResponseOrder struct {
	Response Response      `json:"response"`
	Orders   []model.Order `json:"orders,omitempty"`
	Order    *model.Order  `json:"order,omitempty"`
}

// * Create order
func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// getting data from request
	restaurantId, ok := r.Context().Value("id").(uint)
	if !ok {
		h.logger.Error("can't get id from context")
		SendError(w, "Bad request", 400)
		return
	}
	order := model.Order{}
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		h.logger.Error("Can't decode order", err)
		SendError(w, "Can't get data", http.StatusBadRequest)
		return
	}
	order.RestaurantId = restaurantId

	//validation
	if err := validator.New().Struct(order); err != nil {
		h.logger.Error("validation err: ", err)
		resp := ValidateError(err.(validator.ValidationErrors))
		SendRespError(w, resp, 400)
		return
	}

	// creating order
	if err := h.service.OrderService.Create(order); err != nil {
		h.logger.Error("Can't create order", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}

	// sending response
	resp := Response{Success: true, Message: "Order is created"}
	json.NewEncoder(w).Encode(resp)
}

// * Get all orders
func (h *Handler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// getting data from request
	RId := mux.Vars(r)["restaurant_id"]
	u64, _ := strconv.ParseUint(RId, 10, 32)
	restaurantId := uint(u64)

	// getting orders
	orders, err := h.service.OrderService.GetAll(restaurantId)
	if err != nil {
		h.logger.Error("can't get orders", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}

	// sending response
	resp := ResponseOrder{
		Response: Response{
			Success: true,
		},
		Orders: orders,
	}
	json.NewEncoder(w).Encode(&resp)
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
		h.logger.Error("can't get order", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}

	// sending response
	resp := ResponseOrder{
		Response: Response{
			Success: true,
		},
		Order: order,
	}
	json.NewEncoder(w).Encode(&resp)
}

// * Delete
func (h *Handler) DeleteOrderById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//getting data from request
	OId := mux.Vars(r)["order_id"]
	u64, _ := strconv.ParseUint(OId, 10, 32)
	OrderId := uint(u64)

	// deleting an order
	if err := h.service.OrderService.Delete(OrderId); err != nil {
		h.logger.Error("Can't delete an order", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}

	// sending response
	resp := Response{
		Success: true,
	}
	json.NewEncoder(w).Encode(&resp)
}

// * Change status
func (h *Handler) ChangeOrderStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// getting data from request
	data := struct {
		Status string `json:"status"`
	}{}
	OId := mux.Vars(r)["order_id"]
	u64, _ := strconv.ParseUint(OId, 10, 32)
	OrderId := uint(u64)
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		h.logger.Error("Can't decode data for status", err)
		SendError(w, "Can't get data", http.StatusBadRequest)
		return
	}

	// Validate status
	if !(data.Status == "IN_PROCCESS" || data.Status == "READY" || data.Status == "RECIEVED") {
		h.logger.Error("wrong status form")
		SendError(w, "There are only  3 statuses: `IN_PROCCESS`, `READY` and `RECIEVED`", http.StatusBadRequest)
		return
	}

	// Changing order status
	statusDB, err := h.service.OrderService.ChangeStatus(OrderId, data.Status)
	if err != nil {
		h.logger.Error("Can't delete an order", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}

	// sending response
	resp := Response{
		Success: true,
		Message: "status is " + statusDB,
	}
	json.NewEncoder(w).Encode(&resp)
}
