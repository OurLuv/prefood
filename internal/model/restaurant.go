package model

import "time"

type Restaurant struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name" validate:"required,max=255"`
	ClientId  uint      `json:"client_id"`
	Phone     string    `json:"phone" validate:"required,e164"`
	Country   string    `json:"country"`
	State     string    `json:"state"`
	City      string    `json:"city"`
	Street    string    `json:"street"`
	Email     string    `json:"email" validate:"required,email"`
	Open      bool      `json:"open"`
	CreatedAt time.Time `json:"time"`
}
