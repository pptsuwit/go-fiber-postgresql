package model

import "time"

type Login struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
type LoginWithUser struct {
	Token string `json:"token"`
	User  User
}

type Register struct {
	Username  string `json:"username" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	// Phone     string `json:"phone" validate:"required,len=10,numeric"`
}

type ClaimsToken struct {
	Id  string    `json:"userId"`
	Exp time.Time `json:"exp"`
	Iat time.Time `json:"iat"`
}
