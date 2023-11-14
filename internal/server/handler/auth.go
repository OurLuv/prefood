package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/OurLuv/prefood/internal/server/middleware"
	"github.com/go-playground/validator/v10"
)

// * Login
func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	// getting data from request
	data := struct {
		Email    string `json:"email" validate:"required,email,max=255"`
		Password string `json:"password" validate:"required,max=255"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		h.logger.Error("getting data from request err: ", err)
		SendError(w, "There is no data", http.StatusBadRequest)
		return
	}

	// validate
	if err := validator.New().Struct(data); err != nil {
		h.logger.Error("validation err: ", err)
		resp := ValidateError(err.(validator.ValidationErrors))
		SendRespError(w, resp, 400)
		return
	}

	// looking for user
	user, err := h.service.UserService.Login(data.Email, data.Password)
	if err != nil {
		h.logger.Error("can't find an user: ", err)
		SendError(w, "Email or password is incorrect", http.StatusBadRequest)
		return
	}

	// log user in
	token, err := middleware.CreateToken(user.Id)
	if err != nil {
		h.logger.Error("can't create a token: ", err)
		SendError(w, "Email or password is incorrect", http.StatusBadRequest)
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
	response := Response{
		Success: true,
	}
	json.NewEncoder(w).Encode(response)
}

// * Signup
func (h *Handler) signup(w http.ResponseWriter, r *http.Request) {

	// getting user
	user := model.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.logger.Error("getting data from request err: ", err)
		SendError(w, "Can't get a data", http.StatusBadRequest)
		return
	}

	// validation
	if err := validator.New().Struct(user); err != nil {
		resp := ValidateError(err.(validator.ValidationErrors))
		SendRespError(w, resp, 400)
		return
	}

	// creating user
	if err := h.service.UserService.Create(user); err != nil {
		h.logger.Error("user create error: ", err)
		SendError(w, "can't sign up", 500)
		return
	}
}

// * Signout
func (h *Handler) signout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
}
