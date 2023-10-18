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

	//*Food
	r.HandleFunc("/food", h.GetAllFood).Methods("GET")
	return r
}

func NewHandler(s service.Service) *Handler {
	return &Handler{
		service: s,
	}
}
