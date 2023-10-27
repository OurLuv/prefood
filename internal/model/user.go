package model

type User struct {
	Id        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}
