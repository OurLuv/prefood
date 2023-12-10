package model

type Сategory struct {
	Id           uint   `json:"id"`
	Name         string `json:"name" validate:"required,max=255"`
	RestaurantId uint   `json:"restaurant_id"`
}
