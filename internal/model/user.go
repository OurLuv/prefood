package model

type User struct {
	Id        uint   `json:"id"`
	Firstname string `json:"firstname" validate:"required,max=255"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password" validate:"required,max=255"`
	Email     string `json:"email" validate:"required,email,max=255"`
}
