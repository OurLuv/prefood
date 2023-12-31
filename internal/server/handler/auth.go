package handler

import (
	"encoding/json"
	"net/http"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/OurLuv/prefood/internal/server/middleware"
	"github.com/go-playground/validator/v10"
)

type UserLogin struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,max=255"`
}

type ResponseToken struct {
	Status int    `json:"status"`
	Token  string `json:"token"`
}

// * Login
// @Summary SignIn
// @Tags Auth
// @Description sign in account
// @ID sign-in-account
// @Accept json
// @Produce json
// @Param data body UserLogin true "account info"
// @Success default {object} ResponseToken
// @Failure default {object} Response
// @Router /login [post]
func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// getting data from request

	data := UserLogin{}
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
	// cookie := http.Cookie{
	// 	Name:     "token",
	// 	Value:    token,
	// 	Path:     "/",
	// 	MaxAge:   3600,
	// 	HttpOnly: true,
	// 	Secure:   true,
	// 	SameSite: http.SameSiteLaxMode,
	// }
	//http.SetCookie(w, &cookie)

	response := ResponseToken{
		Status: 1,
		Token:  token,
	}
	json.NewEncoder(w).Encode(response)
}

// * Signup
// @Summary SignUp
// @Tags Auth
// @Description create account
// @ID create-account
// @Accept json
// @Produce json
// @Param input body model.User true "account info"
// @Response default {object} Response
// @Router /signup [post]
func (h *Handler) signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	if err := h.service.UserService.CheckForEmail(user.Email); err != nil {
		if err.Error() == "this email is already in use" {
			SendError(w, "this email is already in use", http.StatusBadRequest)
			return
		}
		h.logger.Error("can't check for an email", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}

	// creating user
	var u *model.User
	var err error
	if u, err = h.service.UserService.Create(user); err != nil {
		h.logger.Error("user create error: ", err)
		SendError(w, "can't sign up", 500)
		return
	}
	u.Password = ""
	response := Response{
		Status: 1,
		Data:   u,
	}

	json.NewEncoder(w).Encode(response)
}

// * Signout
func (h *Handler) signout(w http.ResponseWriter, r *http.Request) {

}
