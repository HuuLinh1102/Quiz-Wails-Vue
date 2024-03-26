package model

type LoginResponse struct {
	Message string
	Status bool
	User *User
}