package model

import "time"

type Food struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	CategoryId  uint      `json:"category-id"`
	Category    Сategory  `json:"category"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	InStock     bool      `json:"in_stock"`
	CreatedAt   time.Time `json:"time"`
	Image       string    `json:"image"`
}
