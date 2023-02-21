package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"unique;not null" json:"username"`
	Password  string
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type CreateUserRequest struct {
	Username        string `json:"username" binding:"required"`
	FirstName       string `json:"first_name" binding:"required"`
	LastName        string `json:"last_name" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}
