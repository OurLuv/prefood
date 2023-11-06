package model

import "time"

type Restaurant struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	ClientId  uint      `json:"client_id"`
	Phone     string    `json:"phone"`
	Country   string    `json:"country"`
	State     string    `json:"state"`
	City      string    `json:"city"`
	Street    string    `json:"street"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"time"`
}
