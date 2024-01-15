package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(50);unique;not null"`
	Password  string `gorm:"type:varchar(200);unique;not null"`
	FirstName string `gorm:"type:varchar(50);not null"`
	LastName  string `gorm:"type:varchar(50);not null"`
	// Phone    string `gorm:"type:varchar(50);unique;not null"`
}
type UserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserRequest struct {
	Username  string `json:"username" validate:"required,email"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Password  string `json:"password" validate:"required,min=6"`
}

type UserResponseWithPagination struct {
	User       []UserResponse     `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}
