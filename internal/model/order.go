package model

import (
	"time"
)

type Order struct {
	Id           uint        `json:"id"`
	RestaurantId uint        `json:"restaurant_id"`
	FoodOrder    []FoodOrder `json:"food_order"`
	Food         []Food
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	Total        int       `json:"total"`
	Status       string    `json:"status"`
	Channel      string    `json:"channel"`
	Additive     string    `json:"additive"`
	Discount     int       `json:"discount"`
	Ordered      time.Time `json:"ordered"`
	Arrive       time.Time `json:"arrive"`
}

type FoodOrder struct {
	Id       uint `json:"id"`
	Quantity int  `json:"quantity"`
}
