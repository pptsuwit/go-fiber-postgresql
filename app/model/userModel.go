package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(50);unique;not null"`
	Password  string `gorm:"type:varchar(200);not null"`
	FirstName string `gorm:"type:varchar(50);not null"`
	LastName  string `gorm:"type:varchar(50);not null"`

	AssetId *uint `gorm:"type:bigint;default:null"`
	Asset   Asset `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; default:null"`
}
type UserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	// AssetId   uint   `json:"assetId"`
	AssetFile string `json:"assetFile"`
}

type UserRequest struct {
	Id        uint   `json:"id; default:null"`
	Username  string `json:"username" validate:"required,email"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Password  string `json:"password" validate:"required,min=6"`

	AssetId       *uint
	AssetHostPath string
}

type UpdateUserRequest struct {
	Id        uint   `json:"id; default:null"`
	Username  string `json:"username" validate:"required,email"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`

	AssetId       *uint
	AssetHostPath string
}

type UserResponseWithPagination struct {
	User       []UserResponse     `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

type UserAsset struct {
	HostName   string
	FolderName string
}
