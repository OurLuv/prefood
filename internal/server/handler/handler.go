package handler

import (
	"github.com/OurLuv/prefood/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	service service.Service
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
	r.HandleFunc("/restaurants/add", h.CreateRestaurant).Methods("POST")
	r.HandleFunc("/restaurants/{id}", h.GetRestaurantById).Methods("GET")

	//*Food
	r.HandleFunc("/menu", h.GetAllFood).Methods("GET")
	r.HandleFunc("/menu/item/{id}", h.GetFoodById).Methods("GET")
	r.HandleFunc("/menu/add", h.CreateFood).Methods("POST")
	r.HandleFunc("/menu/add", h.CreateFoodView).Methods("GET")

	//*Auth
	r.HandleFunc("/login", h.login).Methods("POST")
	r.HandleFunc("/signup", h.signup).Methods("POST")

	return r
}

func NewHandler(s service.Service) *Handler {
	return &Handler{
		service: s,
	}
}
