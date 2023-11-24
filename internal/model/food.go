package model

import "time"

type Food struct {
	Id           uint      `json:"id"`
	RestaurantId uint      `json:"restaurant_id"`
	Name         string    `json:"name" validate:"required,max=255"`
	CategoryId   uint      `json:"category-id"`
	Category     Сategory  `json:"category"`
	Description  string    `json:"description"`
	Price        int       `json:"price" validate:"required"`
	InStock      bool      `json:"in_stock" default:"true"`
	CreatedAt    time.Time `json:"time"`
	Image        string    `json:"image"`
}
