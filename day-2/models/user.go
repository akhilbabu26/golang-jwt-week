package models

type User struct{
	ID uint
	Email string `json:"email"`
	Password string `json:"password"`
}
