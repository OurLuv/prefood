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
	Code         string    `json:"discount"`
	Ordered      time.Time `json:"ordered"`
	Arrive       time.Time `json:"arrive"`
}

type CreateOrderRequest struct {
	Name      string      `json:"name"`
	Phone     string      `json:"phone"`
	Channel   string      `json:"channel"`
	Additive  string      `json:"additive"`
	Code      string      `json:"discount"`
	FoodOrder []FoodOrder `json:"food_order"`
}

type FoodOrder struct {
	Id       uint `json:"id"`
	Quantity int  `json:"quantity"`
}
