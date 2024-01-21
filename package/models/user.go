package models

import "gorm.io/gorm"

type UserDetails struct {
	gorm.Model
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Address string `json:"address"`
}
