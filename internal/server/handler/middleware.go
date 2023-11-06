package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/OurLuv/prefood/internal/server/middleware"
	"github.com/gorilla/mux"
)

// * Check for auth
func (h *Handler) userIdentity(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Not authorized 1: "+err.Error(), http.StatusUnauthorized)
			return
		}
		token := c.Value
		id, err := middleware.VerifyToken(token)
		newCtx := context.WithValue(r.Context(), "id", id)
		if err != nil {
			http.Error(w, "Not authorized 2", http.StatusUnauthorized)
			return
		}
		next(w, r.WithContext(newCtx))
	}
}

// * Check for restautants
func (h *Handler) restaurantAccess(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Not authorized 1: "+err.Error(), http.StatusUnauthorized)
			return
		}
		token := c.Value
		client_id, err := middleware.VerifyToken(token)
		if err != nil {
			http.Error(w, "Not authorized 1: "+err.Error(), http.StatusUnauthorized)
			return
		}

		RId := mux.Vars(r)["id"]
		u64, _ := strconv.ParseUint(RId, 10, 32)
		restaurantId := uint(u64)

		restaurant, err := h.service.RestaruantService.GetById(restaurantId, client_id)
		if err != nil {
			http.Error(w, "Not authorized 3: "+err.Error(), http.StatusUnauthorized)
			return
		}
		newCtx := context.WithValue(r.Context(), "restaurant", restaurant)
		if err != nil {
			http.Error(w, "Not authorized 2", http.StatusUnauthorized)
			return
		}
		next(w, r.WithContext(newCtx))
	}
}
