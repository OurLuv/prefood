package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/OurLuv/prefood/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"
)

type Handler struct {
	service service.Service
	logger  *slog.Logger
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()
	//* Category API
	r.HandleFunc("/api/category/create", h.CreateCategory).Methods("POST")
	r.HandleFunc("/api/category/{id}", h.GetCategoryById).Methods("GET")
	r.HandleFunc("/api/category", h.GetCategories).Methods("GET")
	r.HandleFunc("/api/category/{id}", h.DeleteCategoryById).Methods("DELETE")
	r.HandleFunc("/api/category/{id}", h.UpdateCategoryById).Methods("UPDATE")

	//* Restaruant
	r.HandleFunc("/restaurants", h.userIdentity(h.GetAllRestaurants)).Methods("GET")
	r.HandleFunc("/restaurants/add", h.userIdentity(h.CreateRestaurant)).Methods("POST")
	r.HandleFunc("/restaurants/{id}", h.restaurantAccess(h.GetRestaurantById)).Methods("GET")
	r.HandleFunc("/restaurants/{id}", h.restaurantAccess(h.DeleteRestaurant)).Methods("DELETE")
	r.HandleFunc("/restaurants/{id}/openclose", h.restaurantAccess(h.OpenClose)).Methods("POST")
	r.HandleFunc("/restaurants/{id}", h.restaurantAccess(h.UpdateRestaurant)).Methods("PUT")

	//*Food
	r.HandleFunc("/menu", h.GetAllFood).Methods("GET")
	r.HandleFunc("/menu/item/{id}", h.GetFoodById).Methods("GET")
	r.HandleFunc("/menu/add", h.CreateFood).Methods("POST")
	r.HandleFunc("/menu/add", h.CreateFoodView).Methods("GET")

	//*Auth
	r.HandleFunc("/login", h.login).Methods("POST")
	r.HandleFunc("/signup", h.signup).Methods("POST")
	r.HandleFunc("/signout", h.signout).Methods("GET")

	//*Order
	r.HandleFunc("/restaurants/{restaurant_id}/orders", h.orderAccess(h.CreateOrder)).Methods("POST")
	r.HandleFunc("/restaurants/{restaurant_id}/orders", h.orderAccess(h.GetAllOrders)).Methods("GET")
	r.HandleFunc("/restaurants/{restaurant_id}/orders/{order_id}", h.orderAccess(h.GetOrderById)).Methods("GET")

	return r
}

func NewHandler(s service.Service, l *slog.Logger) *Handler {
	return &Handler{
		service: s,
		logger:  l,
	}
}

func SendError(w http.ResponseWriter, errorStr string, code int) {
	w.WriteHeader(code)
	response := Response{
		Success: false,
		Error:   errorStr,
	}
	json.NewEncoder(w).Encode(response)
}

func SendRespError(w http.ResponseWriter, resp Response, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

func ValidateError(errs validator.ValidationErrors) Response {
	var msgs []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			msgs = append(msgs, fmt.Sprintf("field %s is a required field", err.Field()))
		case "email":
			msgs = append(msgs, fmt.Sprintf("field %s has to be an email", err.Field()))
		case "max":
			msgs = append(msgs, fmt.Sprintf("field %s is wrong length", err.Field()))
		default:
			msgs = append(msgs, fmt.Sprintf("field %s is not valid", err.Field()))
		}
	}
	return Response{
		Success: false,
		Error:   strings.Join(msgs, "; "),
	}
}
