package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/OurLuv/prefood/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type FoodResponse struct {
	Response Response     `json:"response"`
	Menu     []model.Food `json:"menu,omitempty"`
	Food     *model.Food  `json:"food,omitempty"`
}

// * Get all
// @Summary GetFood
// @Tags Menu
// @Description get food
// @ID get-food
// @Param restaurant_id path int true "restaurant id"
// @Produce json
// @Success 200 {object} FoodResponse
// @Failure default {object} Response
// @Router /restaurants/{restaurant_id}/menu [get]
func (h *Handler) GetAllFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//getting data from url
	RId := mux.Vars(r)["restaurant_id"]
	u64, _ := strconv.ParseUint(RId, 10, 32)
	restaurantId := uint(u64)

	// looking for all restaurants
	var food []model.Food
	var err error
	if food, err = h.service.FoodService.GetAll(restaurantId); err != nil {
		h.logger.Error("storage error: ", err)
		SendError(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	// sending response
	resp := FoodResponse{
		Response: Response{Status: 1},
		Menu:     food,
	}
	json.NewEncoder(w).Encode(resp)

}

// * Get food by id
// @Summary GetFoodById
// @Tags Menu
// @Description get food by id
// @ID get-food-by-id
// @Param restaurant_id path int true "restaurant id"
// @Param id path int true "food id"
// @Produce json
// @Success 200 {object} FoodResponse
// @Failure default {object} Response
// @Router /restaurants/{restaurant_id}/menu/{id} [get]
func (h *Handler) GetFoodById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//getting data from url
	RId := mux.Vars(r)["restaurant_id"]
	u64, _ := strconv.ParseUint(RId, 10, 32)
	restaurantId := uint(u64)
	FId := mux.Vars(r)["id"]
	u64, _ = strconv.ParseUint(FId, 10, 32)
	foodId := uint(u64)

	// looking for all restaurants
	var food *model.Food
	var err error
	if food, err = h.service.FoodService.GetById(restaurantId, foodId); err != nil {
		h.logger.Error("storage error: ", err)
		SendError(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	// sending response
	resp := FoodResponse{
		Response: Response{Status: 1},
		Food:     food,
	}
	json.NewEncoder(w).Encode(resp)
}

// * Create food
// @Summary CreateFood
// @Security ApiKeyAuth
// @Tags Menu
// @Description create food
// @ID create-food
// @Param restaurant_id path int true "restaurant id"
// @Accept mpfd
// @Produce json
// @Param image formData file true "food image"
// @Param food formData string true "food info"
// @Success 200 {object} FoodResponse
// @Failure default {object} Response
// @Router /restaurants/{restaurant_id}/menu [post]
func (h *Handler) CreateFood(w http.ResponseWriter, r *http.Request) {
	var food model.Food
	var err error

	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		h.logger.Error("can't parse form", err)
		SendError(w, "Bad request", http.StatusBadRequest)
		return
	}

	// getting data from request
	jsonData := r.FormValue("food")
	if err := json.Unmarshal([]byte(jsonData), &food); err != nil {
		h.logger.Error("can't get data from request", err)
		SendError(w, "There is no data", http.StatusBadRequest)
		return
	}
	file, header, err := r.FormFile("image")
	if err != nil {
		h.logger.Error("can't upload image", err)
		SendError(w, "Error with uploading an image", http.StatusBadRequest)
		return
	}
	imgName := strings.Split(header.Filename, ".")
	imgType := imgName[len(imgName)-1]
	food.Image = imgType

	// setting restaurant id from context
	restaurant, ok := r.Context().Value("restaurant").(*model.Restaurant)
	if !ok {
		h.logger.Error("Can't get restaurant from context")
		SendError(w, "Can't get a restauarant", http.StatusInternalServerError)
		return
	}
	food.RestaurantId = restaurant.Id
	if food.Category.Name == "" {
		food.Category.Name = "Not Empty"
	}
	// validatation
	if err := validator.New().Struct(food); err != nil {
		h.logger.Error("validation err: ", err)
		resp := ValidateError(err.(validator.ValidationErrors))
		SendRespError(w, resp, 400)
		return
	}
	if food.Image != "jpg" && food.Image != "png" && food.Image != "jpeg" {
		SendError(w, "Image has to be jpg/jpeg or png", http.StatusBadRequest)
		return
	}

	// creating food
	f, err := h.service.FoodService.Create(food)
	if err != nil {
		h.logger.Error("can't create food", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}
	//downloading image on server
	fileContent, err := io.ReadAll(file)
	if err != nil {
		h.logger.Error("can't read a file", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}
	newFile, err := os.Create("static/images/" + f.Image)
	if err != nil {
		h.logger.Error("can't create a file", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = newFile.Write(fileContent)
	if err != nil {
		h.logger.Error("can't add a content to the file", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}

	// sending response
	resp := FoodResponse{
		Response: Response{Status: 1, Message: "Row is added to database"},
		Food:     f,
	}
	json.NewEncoder(w).Encode(resp)
}

// * Update
// @Summary UpdateFood
// @Security ApiKeyAuth
// @Tags Menu
// @Description update food
// @ID update-food
// @Param restaurant_id path int true "restaurant id"
// @Accept mpfd
// @Produce json
// @Param image formData file true "food image"
// @Param food formData string true "food info"
// @Success 200 {object} FoodResponse
// @Failure default {object} Response
// @Router /restaurants/{restaurant_id}/menu [put]
func (h *Handler) UpdateFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var food model.Food
	var err error

	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		h.logger.Error("can't parse form", err)
		SendError(w, "Bad request", http.StatusBadRequest)
		return
	}

	// getting data from request
	jsonData := r.FormValue("food")
	if err := json.Unmarshal([]byte(jsonData), &food); err != nil {
		h.logger.Error("can't get data from request", err)
		SendError(w, "There is no data", http.StatusBadRequest)
		return
	}
	file, header, err := r.FormFile("image")
	if err != nil {
		h.logger.Error("can't upload image", err)
		SendError(w, "Error with uploading an image", http.StatusBadRequest)
		return
	}
	imgName := strings.Split(header.Filename, ".")
	imgType := imgName[len(imgName)-1]
	food.Image = imgType

	// setting restaurant id from context
	restaurant, ok := r.Context().Value("restaurant").(*model.Restaurant)
	if !ok {
		h.logger.Error("Can't get restaurant from context")
		SendError(w, "Can't get a restauarant", http.StatusInternalServerError)
		return
	}
	food.RestaurantId = restaurant.Id

	// validatation
	if err := validator.New().Struct(food); err != nil {
		h.logger.Error("validation err: ", err)
		resp := ValidateError(err.(validator.ValidationErrors))
		SendRespError(w, resp, 400)
		return
	}
	if food.Image != "jpg" && food.Image != "png" && food.Image != "jpeg" {
		SendError(w, "Image has to be jpg/jpeg or png", http.StatusBadRequest)
		return
	}

	// updating food
	f, err := h.service.FoodService.UpdateById(food)
	if err != nil {
		h.logger.Error("can't update food", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}
	//downloading image on server
	fileContent, err := io.ReadAll(file)
	if err != nil {
		h.logger.Error("can't read a file", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}
	newFile, err := os.Create("static/images/" + f.Image)
	if err != nil {
		h.logger.Error("can't create a file", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = newFile.Write(fileContent)
	if err != nil {
		h.logger.Error("can't add a content to the file", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}
	resp := Response{
		Status:  1,
		Message: "Item is updated",
	}
	json.NewEncoder(w).Encode(resp)
}

// * Delete
func (h *Handler) DeleteFood(w http.ResponseWriter, r *http.Request) {

}

// func (h *Handler) CreateFoodView(w http.ResponseWriter, r *http.Request) {
// 	t, err := template.ParseFiles("static/create-food.html")
// 	if err != nil {
// 		errStr := fmt.Sprintf("can't load a view 01: %s", err.Error())
// 		http.Error(w, errStr, http.StatusBadRequest)
// 		return
// 	}
// 	mp := make(map[string][]model.Сategory)
// 	if mp["Categories"], err = h.service.CategoryService.GetAll(); err != nil {
// 		fmt.Print(err.Error())
// 	}

// 	if err = t.Execute(w, mp); err != nil {
// 		errStr := fmt.Sprintf("can't load a view 02: %s", err.Error())
// 		http.Error(w, errStr, http.StatusBadRequest)
// 		return
// 	}
// }
