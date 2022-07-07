package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	User_Name string `gorm:"type:varchar(25);unique;not null" json:"user_name" form:"user_name"`
	Name      string `gorm:"type:varchar(255)" json:"name" form:"name"`
	Email     string `gorm:"type:varchar(100);unique;not null" json:"email" form:"email"`
	Password  string `gorm:"type:varchar(255);not null" json:"password" form:"password"`
	Role      string `gorm:"type:varchar(255);not null" json:"role" form:"role"`
	Token     string `json:"token" form:"token"`
}

type UserLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
