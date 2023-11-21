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
		w.Header().Set("Content-Type", "application/json")
		c, err := r.Cookie("token")
		if err != nil {
			h.logger.Error("can't get cookie: ", err)
			SendError(w, "Not authorized", http.StatusUnauthorized)
			return
		}
		token := c.Value
		id, err := middleware.VerifyToken(token)
		if err != nil {
			h.logger.Error("verifying error", err)
			SendError(w, "Not authorized", http.StatusMethodNotAllowed)
			return
		}
		newCtx := context.WithValue(r.Context(), "id", id)
		next(w, r.WithContext(newCtx))
	}
}

// * Check for restautants
func (h *Handler) restaurantAccess(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		c, err := r.Cookie("token")
		if err != nil {
			h.logger.Error("can't get cookie: ", err)
			SendError(w, "Not authorized", http.StatusUnauthorized)
			return
		}
		token := c.Value
		client_id, err := middleware.VerifyToken(token)
		if err != nil {
			h.logger.Error("verifying error", err)
			SendError(w, "Not authorized", http.StatusMethodNotAllowed)
			return
		}

		RId := mux.Vars(r)["restaurant_id"]
		u64, _ := strconv.ParseUint(RId, 10, 32)
		restaurantId := uint(u64)

		restaurant, err := h.service.RestaruantService.GetById(restaurantId, client_id)
		if err != nil {
			h.logger.Error("can't get restaurant", err)
			SendError(w, "Not allowed", http.StatusMethodNotAllowed)
			return
		}
		newCtx := context.WithValue(r.Context(), "restaurant", restaurant)
		next(w, r.WithContext(newCtx))
	}
}

func (h *Handler) orderAccess(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		c, err := r.Cookie("token")
		if err != nil {
			h.logger.Error("can't get cookie: ", err)
			SendError(w, "Not authorized", http.StatusUnauthorized)
			return
		}
		token := c.Value
		client_id, err := middleware.VerifyToken(token)
		if err != nil {
			h.logger.Error("verifying error", err)
			SendError(w, "Not authorized", http.StatusMethodNotAllowed)
			return
		}
		RId := mux.Vars(r)["restaurant_id"]
		u64, _ := strconv.ParseUint(RId, 10, 32)
		restaurantId := uint(u64)
		// check if client has access to this restaurant
		_, err = h.service.RestaruantService.GetById(restaurantId, client_id)
		if err != nil {
			h.logger.Error("can't get restaurant", err)
			SendError(w, "Not allowed", http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	}
}
