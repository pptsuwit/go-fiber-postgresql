package model

import (
	"time"

	"gorm.io/gorm"
)

type CustomerResponseWithPagination struct {
	Customer   []CustomerResponse `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}
type CustomerResponse struct {
	ID         uint   `json:"id"`
	CustomerId uint   `json:"customerId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Address    string `json:"address"`
	Birthdate  string `json:"birthdate"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`

	// GenderId uint   `json:"genderId"`
	// Gender   Gender `json:"genders"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type CustomerRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Birthdate string `json:"birthdate" validate:"required"`
	Address   string `json:"address" `
	Phone     string `json:"phone" validate:"required"`
	Email     string `json:"email" validate:"required" valid:"email"`

	// GenderId uint `json:"genderId" validate:"required"`
}

type Customer struct {
	gorm.Model
	CustomerId uint   `gorm:"type:uint;not null"`
	FirstName  string `gorm:"type:varchar(50);not null"`
	LastName   string `gorm:"type:varchar(50);not null"`
	Birthdate  string `gorm:"type:varchar(200);not null"`
	Address    string `gorm:"type:varchar(200);"`
	Phone      string `gorm:"type:varchar(50);not null"`
	Email      string `gorm:"type:varchar(50);unique;not null"`
}

type CustomerCreate struct {
	CustomerId uint   `json:"customerId" validate:"required"`
	FirstName  string `json:"firstName" validate:"required"`
	LastName   string `json:"lastName" validate:"required"`
	Birthdate  string `json:"birthdate" validate:"required"`
	Address    string `json:"address" `
	Phone      string `json:"phone" validate:"required"`
	Email      string `json:"email" validate:"required" valid:"email"`

	// GenderId uint `json:"genderId" validate:"required"`
}
