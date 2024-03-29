package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "reflect"
	"strings"

	_ "github.com/OurLuv/prefood/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/OurLuv/prefood/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"golang.org/x/exp/slog"
)

type Handler struct {
	service service.Service
	logger  *slog.Logger
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type ResponseData struct {
	Response
	Data interface{}
}

type ResponseId struct {
	Response Response `json:"response,omitempty"`
	Id       uint     `json:"id,omitempty"`
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	//* Category API
	r.HandleFunc("/restaurants/{restaurant_id}/category", h.restaurantAccess(h.CreateCategory)).Methods("POST")
	r.HandleFunc("/restaurants/{restaurant_id}/category/{category_id}", h.GetCategoryById).Methods("GET")
	r.HandleFunc("/restaurants/{restaurant_id}/category", h.GetCategories).Methods("GET")
	r.HandleFunc("/restaurants/{restaurant_id}/category/{category_id}", h.restaurantAccess(h.DeleteCategoryById)).Methods("DELETE")
	// todo: if 0 rows werer Update - particular message
	r.HandleFunc("/restaurants/{restaurant_id}/category/{category_id}", h.restaurantAccess(h.UpdateCategoryById)).Methods("PUT")

	//* Restaruant
	r.HandleFunc("/restaurants", h.userIdentity(h.GetAllRestaurants)).Methods("GET")
	r.HandleFunc("/restaurants", h.userIdentity(h.CreateRestaurant)).Methods("POST")
	r.HandleFunc("/restaurants/{restaurant_id}", h.restaurantAccess(h.GetRestaurantById)).Methods("GET")
	r.HandleFunc("/restaurants/{restaurant_id}", h.restaurantAccess(h.DeleteRestaurant)).Methods("DELETE")
	r.HandleFunc("/restaurants/{restaurant_id}/openclose", h.restaurantAccess(h.OpenClose)).Methods("POST")
	r.HandleFunc("/restaurants/{restaurant_id}", h.restaurantAccess(h.UpdateRestaurant)).Methods("PUT")

	//*Food
	// todo: fix a bug with categories
	r.HandleFunc("/restaurants/{restaurant_id}/menu", h.GetAllFood).Methods("GET")
	r.HandleFunc("/restaurants/{restaurant_id}/menu/{id}", h.GetFoodById).Methods("GET")
	r.HandleFunc("/restaurants/{restaurant_id}/menu", h.restaurantAccess(h.CreateFood)).Methods("POST")
	// todo: if 0 rows werer Update - particular message
	r.HandleFunc("/restaurants/{restaurant_id}/menu/{id}", h.restaurantAccess(h.UpdateFood)).Methods("PUT")

	//*Auth
	r.HandleFunc("/login", h.login).Methods("POST")
	r.HandleFunc("/signup", h.signup).Methods("POST")
	//r.HandleFunc("/signout", h.signout).Methods("GET")

	//*Order
	r.HandleFunc("/restaurants/{restaurant_id}/orders", h.orderAccess(h.CreateOrder)).Methods("POST")
	r.HandleFunc("/restaurants/{restaurant_id}/orders", h.orderAccess(h.GetAllOrders)).Methods("GET")
	r.HandleFunc("/restaurants/{restaurant_id}/orders/{order_id}", h.orderAccess(h.GetOrderById)).Methods("GET")
	r.HandleFunc("/restaurants/{restaurant_id}/orders/{order_id}", h.orderAccess(h.DeleteOrderById)).Methods("DELETE")
	r.HandleFunc("/restaurants/{restaurant_id}/orders/{order_id}", h.orderAccess(h.ChangeOrderStatus)).Methods("POST")

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
		Status: -1,
		Error:  errorStr,
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
		Status: -1,
		Error:  strings.Join(msgs, "; "),
	}
}
