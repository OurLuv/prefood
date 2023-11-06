package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/OurLuv/prefood/internal/server/middleware"
)

// * Login
func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := h.service.UserService.Login(data.Email, data.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token, err := middleware.CreateToken(user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
}

// * Signup
func (h *Handler) signup(w http.ResponseWriter, r *http.Request) {
	user := model.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.service.UserService.Create(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

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
